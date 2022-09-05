package scheduler

import "crawler/concurrent/engine"

/**
==> 实现简单调度器
*/

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// scheduler 将任务分发到 work chan 中，多个 worker 从 worker chan 中拿 request 进行处理，处理完成的结果包含两部分，
	// ==> 一部分是 item，也就是我们需要的用户信息，现在对其处理的方式是直接打印在控制台上
	// ==> 另一部分是包含在结果中的 request，该 request 会继续被 scheduler 分发到 work chan 中进行处理
	// 此处就出现了【循环引用】的问题：worker 等待 scheduler 分发其处理结果中包含的 request，而 scheduler 等待 worker 接收其分发的任务
	// ==> 解决方案：为每个 request 创建一个 go routine，每个 go routine 只做一件事，就是将 request 分发到 worker chan 这一操作
	go func() {
		// 将请求发送到 work chan 中
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	s.workerChan = requests
}
