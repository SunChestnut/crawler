package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func serverRpc(host, index string) error {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://192.168.56.10:9200/"},
	})
	if err != nil {
		log.Printf("error create elasticsearch client : %v", err)
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}

func main() {
	err := serverRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticSearchIndex)
	if err != nil {
		log.Fatalf("Start rpc server fail : %v", err)
	}
}
