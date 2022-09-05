package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/queuedengine"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
)

func main() {

	e := queuedengine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
