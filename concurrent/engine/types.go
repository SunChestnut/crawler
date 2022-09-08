package engine

// ParserFunc ==> 根据 内容 和 url 即可具有解析功能的函数
type ParserFunc func(contents []byte, url string) ParserResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
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

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
