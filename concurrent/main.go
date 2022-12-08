package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/engine/queue"
	"crawler/concurrent/persist"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
	"crawler/distributed/config"
)

func main() {

	// 如果连接 ElasticSearch 失败，即使项目运行正常但不保存数据也毫无意义，因此终止程序运行
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := queue.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
