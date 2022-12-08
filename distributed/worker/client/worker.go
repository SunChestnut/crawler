package client

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/engine/queue"
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"log"
	"net/rpc"
)

// CreateProcessor ==>
func CreateProcessor(clientChan chan *rpc.Client) queue.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	callWorkerServer := func(request engine.Request) (engine.ParserResult, error) {
		// 序列化 Request，以便在网络中传输
		sRequest := worker.SerializeRequest(request)

		// 调用 worker 服务端
		var sResult worker.ParserResult
		rpcClient := <-clientChan
		// rpc client
		err := rpcClient.Call(config.CrawlServiceRpc, sRequest, &sResult)
		if err != nil {
			return engine.ParserResult{}, err
		}

		return worker.DeserializeParserResult(sResult), nil
	}

	return callWorkerServer
}

// CreateClientPool ==> 根据给定的 hosts 创建 rpc 客户端
func CreateClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("Error connecting to %s: %v", h, err)
			continue
		}
		clients = append(clients, client)
		log.Printf("Worker client connect to %s", h)
	}

	// array to channel
	out := make(chan *rpc.Client)
	go func() {
		// 轮流分发 client 且 每轮分发结束后继续下一轮的分发
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
