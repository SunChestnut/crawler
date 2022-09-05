package parser

import (
	"crawler/stand_alone/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^>]*)</a>`

// ParseCityList 城市列表解析器
func ParseCityList(contents []byte) engine.ParserResult {

	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range all {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
