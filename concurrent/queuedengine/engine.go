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
	WorkerChan() chan engine.Request // 向调度器询问：我有一个 worker，给我哪个 channel
	Run()
	ReadyNotify // 如果不将 ReadyNotify 放入 Scheduler 中，在 Run 函数中调用 createWorker 函数时则会报错
}

// ReadyNotify 在 createWorker 函数中需要使用到 WorkerReady 函数的功能，但在参数中将 Scheduler 全部传入过于繁重，因此将该功能单独提取出来
type ReadyNotify interface {
	WorkerReady(chan engine.Request)
}

func (e *ConcurrentEngine) Run(seeds ...engine.Request) {
	e.Scheduler.Run()
	out := make(chan engine.ParserResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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

func createWorker(in chan engine.Request, out chan engine.ParserResult, ready ReadyNotify) {
	go func() {
		for {
			// 告诉 Scheduler 我已经就绪了，就绪后才能继续接收数据
			ready.WorkerReady(in)
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
