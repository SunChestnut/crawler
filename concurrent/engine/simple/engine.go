package simple

import (
	"crawler/concurrent/engine"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(engine.Request)
	ConfigureMasterWorkerChan(chan engine.Request)
}

func (e *ConcurrentEngine) Run(seeds ...engine.Request) {
	in := make(chan engine.Request)
	out := make(chan engine.ParserResult)

	// 将 Engine 接收到的 request 送入 scheduler 中进行分发处理
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 将从 Worker 中接收的数据分为两部分：打印 item 和 将 request 送入 scheduler 中
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan engine.Request, out chan engine.ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := engine.Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
