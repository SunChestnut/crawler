package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"log"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(500 * time.Millisecond)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Errorf("error start rpc client: %v", err)
	}

	req := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481",
		Parser: worker.SerializedParser{
			FunctionName: config.ParseProfile,
			Args:         "一身傲气如你*",
		},
	}

	var result worker.ParserResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("error call method: %v", err)
	}
	log.Printf("%+v", result)
}
