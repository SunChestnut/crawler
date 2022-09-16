package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/queuedengine"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
	"crawler/distributed/config"
	itemSaver "crawler/distributed/persist/client"
	worker "crawler/distributed/worker/client"
	"fmt"
	"log"
)

func main() {
	// host = :1234
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		log.Fatalf("error start worker client : %v", err)
	}

	e := queuedengine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
