package main

import (
	"context"
	"crawler/config"
	"crawler/engine"
	"crawler/engine/scheduler"
	"crawler/model"
	"crawler/persist/client"
	"crawler/support/grpc"
	wclient "crawler/worker/client"
	zparser "crawler/zhenai/parser"
	"fmt"
	"log"
)

func main() {
	// 启动 ItemSaver 客户端
	itemSaverClient, err := client.StartItemSaverClient(fmt.Sprintf("127.0.0.1:%d", config.ItemSaverPort))
	if err != nil {
		log.Fatalf("💔error start itemSaver client: %v", err)
	}

	// Worker 服务器启动地址
	hosts := []string{
		fmt.Sprintf("127.0.0.1:%d", config.WorkerPort0),
		//fmt.Sprintf("127.0.0.1:%d", config.WorkerPort1),
		//fmt.Sprintf("127.0.0.1:%d", config.WorkerPort2),
	}

	ctx := context.Background()

	workerClientPool := grpc.CreateWorkerClientPool(hosts)
	processor := wclient.CreateProcessor(ctx, workerClientPool)

	e := engine.Engine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      5,
		ItemChan:         itemSaverClient,
		RequestProcessor: processor,
	}
	e.Run(model.Request{
		Url:    config.MockServerUrl,
		Parser: model.NewFuncParser(zparser.ParseCityList, config.ParseCityList),
	})
}
