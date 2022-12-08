package wclient

import (
	"context"
	"crawler/model"
	"crawler/pb"
	"crawler/worker/parser"
	"fmt"
)

func CreateProcessor(clientChan chan pb.CrawlServiceClient) model.Processor {
	return func(request model.Request) (model.ParserResult, error) {
		sRequest := parser.SerializeRequest(request)

		grpcClient := <-clientChan
		parserResult, err := grpcClient.Process(context.Background(), sRequest)
		if err != nil {
			return model.ParserResult{}, fmt.Errorf("[worker.client.CreateProcessor] error call worker server: %v", err)
		}

		result := parser.DeserializeParserResult(parserResult)
		return result, nil
	}
}
