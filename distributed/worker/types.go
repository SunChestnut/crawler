package worker

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/zhenai/parser"
	"crawler/distributed/config"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	FunctionName string
	Args         any
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParserResult struct {
	Items   []engine.Item
	Request []Request
}

/*序列化*/

// SerializeRequest 将 engine.Request 转换成 Request，以便在网络中传输。即序列化 engine.Request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			FunctionName: name,
			Args:         args,
		},
	}
}

// SerializeParserResult 将 engine.ParserResult 转换成 ParserResult，以便在网络中传输。即序列化 engine.ParserRequest
func SerializeParserResult(r engine.ParserResult) ParserResult {
	var reqs []Request
	for _, req := range r.Requests {
		reqs = append(reqs, SerializeRequest(req))
	}
	return ParserResult{
		Items:   r.Items,
		Request: reqs,
	}
}

/*反序列化*/

// DeserializeRequest 将 Request 转换成 engine.Request
func DeserializeRequest(r Request) (engine.Request, error) {
	dParser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: dParser,
	}, nil
}

// deserializeParser 根据 函数名 ➕ 参数 获取对应的函数
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	var ep *engine.FuncParser
	switch p.FunctionName {
	case config.ParseCityList:
		ep = engine.NewFuncParser(parser.ParseCityList, config.ParseCityList)
	case config.ParseCity:
		ep = engine.NewFuncParser(parser.ParseCity, config.ParseCity)
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
	return ep, nil
}

// DeserializeParserResult 将 ParserResult 转换成 engine.ParserResult
func DeserializeParserResult(r ParserResult) engine.ParserResult {
	var reqs []engine.Request
	for _, req := range r.Request {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v\n", err)
			continue
		}
		reqs = append(reqs, engineRequest)
	}
	return engine.ParserResult{
		Requests: reqs,
		Items:    r.Items,
	}
}
