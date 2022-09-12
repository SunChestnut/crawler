package controller

import (
	"context"
	"crawler/concurrent/engine"
	"crawler/frontend/model"
	"crawler/frontend/view"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elasticsearch.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"http://192.168.56.10:9200/"},
		})
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

// url 格式： http://localhost:8888/search?q=男 已购房&from=10
func (h SearchResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 获取 url 中参数 q 后面的内容
	query := strings.TrimSpace(request.FormValue("q"))
	// 获取 url 中的分页值，也就是放在 from 后面的数值
	from, err := strconv.Atoi(request.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(query, from)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(writer, page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(query string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	es := h.client
	resp, err := es.Search(
		es.Search.WithIndex("dating_profile"),
		es.Search.WithQuery(query),
		es.Search.WithFrom(from),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		return result, err
	}

	var r map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %v", err)
		return result, err
	}

	var items []engine.Item
	for _, hit := range r["hits"].(map[string]any)["hits"].([]any) {
		before := hit.(map[string]any)["_source"].(map[string]any)
		items = append(items, convert(before))
	}

	result.Query = query
	result.Hits = r["hits"].(map[string]any)["total"].(map[string]any)["value"].(float64)
	result.Start = from
	result.Items = items
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func convert(before map[string]any) engine.Item {
	marshal, err := json.Marshal(before)
	if err != nil {
		log.Fatalf("Error convert : %v", err)
	}

	var item engine.Item
	err = json.Unmarshal(marshal, &item)
	if err != nil {
		log.Fatalf("Error convert : %v", err)
	}

	return item
}
