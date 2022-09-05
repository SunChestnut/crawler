package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []any
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
