package wclient

import (
	"crawler/model"
	"crawler/pb"
	"crawler/worker/parser"
	"log"
)

func CreateProcessor(clientChan chan *pb.CrawlServiceClient) model.Processor {

	return func(request model.Request) (model.ParserResult, error) {
		log.Printf("[worker.client.CreateProcessor]...")
		sRequest := parser.SerializeRequest(request)

		log.Printf(sRequest.Url)
		//var c *pb.CrawlServiceClient
		//c = <-clientChan
		//grpcClient := <-clientChan
		//parserResult, err := c.Process(context.Background(), sRequest)
		//if err != nil {
		//	return model.ParserResult{}, fmt.Errorf("[worker.client.CreateProcessor] error call worker server: %v", err)
		//}
		//
		//result := parser.DeserializeParserResult(parserResult)
		//return result, nil
		return model.ParserResult{}, nil
	}
}
