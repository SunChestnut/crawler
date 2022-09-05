package parser

import (
	"crawler/stand_alone/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// ParseCity 城市解析器
func ParseCity(contents []byte) engine.ParserResult {

	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}

	for _, m := range all {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParserResult {
				return ParseProfile(contents, name)
			},
		})
	}
	return result
}
