package engine

import (
	"crawler/stand_alone/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := e.worker(r)
		// 在爬虫中，针对某个请求拿不到正确的响应结果是常有的事，为不影响继续处理后续的请求，遇到错误时候略过即可
		if err != nil {
			continue
		}

		// 将结果中的请求再次加入到处理队列中
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// 将 Fetcher 和 Parser 中的功能提取出来形成 worker
func (SimpleEngine) worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s, %v", r.Url, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
