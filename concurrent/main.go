package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"crawler/concurrent/queuedengine"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
)

func main() {

	// 如果连接 ElasticSearch 失败，即使项目运行正常但不保存数据也毫无意义，因此终止程序运行
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := queuedengine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

}
