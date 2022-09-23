package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/queuedengine"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
	"crawler/distributed/config"
	itemSaver "crawler/distributed/persist/client"
	worker "crawler/distributed/worker/client"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	startBasedOnCommand()
}

// startBasedOnConfig => 使用配置的端口启动
func startBasedOnConfig() {
	// host = :1234
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		log.Fatalf("💔 error create itemchan: %v", err)
	}

	hosts := []string{fmt.Sprintf(":%d", config.WorkerPort0), fmt.Sprintf(":%d", config.WorkerPort1)}

	// create client pool
	pool := worker.CreateClientPool(hosts)
	processor := worker.CreateProcessor(pool)

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

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

// startBasedOnCommand => 使用从命令行传入的端口启动
func startBasedOnCommand() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		log.Fatalf("💔error create itemchan: %v", err)
	}

	// create client pool
	pool := worker.CreateClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

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
