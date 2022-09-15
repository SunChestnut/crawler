package engine

import "crawler/distributed/config"

type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args any)
}

type Request struct {
	Url    string
	Parser Parser
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload any
}

type NilParser struct{}

func (NilParser) Parse(contents []byte, url string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args any) {
	return config.NilParser, nil
}

// ParserFunc ==> 根据 内容 和 url 即可具有解析功能的函数
type ParserFunc func(contents []byte, url string) ParserResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

func (p *FuncParser) Parse(contents []byte, url string) ParserResult {
	return p.parser(contents, url)
}

func (p *FuncParser) Serialize() (name string, args any) {
	return p.name, nil
}
