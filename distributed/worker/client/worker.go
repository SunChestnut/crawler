package client

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/queuedengine"
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
)

func CreateProcessor() (queuedengine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (engine.ParserResult, error) {
		sRequest := worker.SerializeRequest(request)
		var sResult worker.ParserResult
		err := client.Call(config.CrawlServiceRpc, sRequest, &sResult)
		if err != nil {
			return engine.ParserResult{}, nil
		}
		return worker.DeserializeParserResult(sResult), nil
	}, nil
}
