package worker

import "crawler/concurrent/engine"

// CrawlService ==> 爬虫服务
type CrawlService struct{}

// Process ==>
func (CrawlService) Process(req Request, result *ParserResult) error {
	// 解析 Request 请求
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParserResult(engineResult)
	return nil
}
