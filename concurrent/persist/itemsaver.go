package persist

import (
	"bytes"
	"context"
	"crawler/concurrent/engine"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func ItemSaver(index string) (chan engine.Item, error) {

	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.56.10:9200/"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("【Item Saver】Got item #%d: %v", itemCount, item)
			itemCount++

			if err := save(client, item, index); err != nil {
				log.Printf("【Item Saver】error saving item : %v", err)
			}
		}
	}()
	return out, nil
}

func save(client *elasticsearch.Client, item engine.Item, index string) error {

	data, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("An error Occured %v", err)
	}

	req := esapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}
	// 当获取的 用户ID 不为空时，使用 用户ID 作为 DocumentID
	if item.Id != "" {
		req.DocumentID = item.Id
	}

	// 发送请求
	resp, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("An error Occured %v", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		log.Printf("[%s] Error", resp.Status())
		return err
	}

	// 解析响应结果并打印
	var r map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return err
	} else {
		//log.Printf("[%s] %s; version=%d", resp.Status(), r["result"], int(r["_version"].(float64)))
	}
	return nil
}

func search(client *elasticsearch.Client, index, id string) engine.Item {

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]any{
		"query": map[string]any{
			"match": map[string]any{
				"ids": []string{id},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(index),
		//client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the response status, number of results, and request duration.
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]any)["total"].(map[string]any)["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	var actual engine.Item
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]any)["hits"].([]any) {

		log.Printf(" * ID=%s, %s", hit.(map[string]any)["_id"], hit.(map[string]any)["_source"])

		receiveId := hit.(map[string]any)["_id"]
		if receiveId == id {
			source := hit.(map[string]any)["_source"]
			actual = convertToProfile(source.(map[string]any))
			log.Printf("Result : %v\n", actual)
		}
	}

	log.Println(strings.Repeat("=", 37))
	return actual
}

func convertToProfile(param map[string]any) (profile engine.Item) {
	// convert map to json
	jsonStr, err := json.Marshal(param)
	log.Printf("JSON-Str: %v\n", jsonStr)
	if err != nil {
		log.Fatalf("Marshal Fail : %v", err)
	}
	if err := json.Unmarshal(jsonStr, &profile); err != nil {
		log.Fatalf("Unmarshal Fail : %v", err)
	}
	return
}
