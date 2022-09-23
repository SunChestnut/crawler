package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	// 使用从命令行传入的端口启动
	flag.Parse()
	if *port == 0 {
		fmt.Println("🙀must specify a port")
		return
	}
	err := serverRpc(fmt.Sprintf(":%d", *port), config.ElasticSearchIndex)

	// 使用配置端口启动
	//err := serverRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticSearchIndex)

	if err != nil {
		log.Fatalf("Start rpc server fail : %v", err)
	}
}

func serverRpc(host, index string) error {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{config.ElasticSearchAddr},
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
