package scheduler

import "crawler/concurrent/engine"

/**
==> 使用队列实现调度器
*/

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request // 存放 worker 的 channel，即每个 worker 都有自己的 channel
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// WorkerReady 告诉外界有 worker 已经就绪了，可以继续接收任务了
func (s *QueuedScheduler) WorkerReady(requests chan engine.Request) {
	s.workerChan <- requests
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 当既有 request 在排队，又有 worker 在排队时，便可以将 request 发给 worker
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
