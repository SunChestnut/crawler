package queuedengine

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/fetcher"
	"log"
)

// ConcurrentEngine 针对【用队列实现调度器】所适配的 Engine/**

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(engine.Request)
	ConfigureMasterWorkerChan(chan engine.Request)
	WorkerReady(chan engine.Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...engine.Request) {
	e.Scheduler.Run()
	out := make(chan engine.ParserResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 将从 worker 中接收的数据分为两部分：打印 item 和 将 request 送入 scheduler 中
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

func createWorker(out chan engine.ParserResult, s Scheduler) {
	in := make(chan engine.Request)
	go func() {
		for {
			// 告诉 Scheduler 我已经就绪了，就绪后才能继续接收数据
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r engine.Request) (engine.ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s, %v", r.Url, err)
		return engine.ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
