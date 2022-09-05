package main

import (
	"crawler/stand_alone/engine"
	"crawler/stand_alone/zhenai/parser"
)

func main() {

	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//resp, err := http.Get("https://album.zhenai.com/u/1318618870")
	//if err != nil {
	//	panic(err)
	//}
	//
	//response, err := httputil.DumpResponse(resp, true)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ioutil.WriteFile("profile_test_data.html", response, 0777)
	//if err != nil {
	//	panic(err)
	//}

	//contents, err := ioutil.ReadFile("./stand_alone/zhenai/test/citylist_test_data.html")
	//if err != nil {
	//	panic(err)
	//}
	//cityList := parser.ParseCityList(contents)
	//fmt.Println(cityList)

}
