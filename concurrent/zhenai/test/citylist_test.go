package test

import (
	"crawler/stand_alone/fetcher"
	"crawler/stand_alone/zhenai/parser"
	"os"
	"testing"
)

func TestFetch(t *testing.T) {

	//const url = "http://www.zhenai.com/zhenghun/"
	const url = "https://album.zhenai.com/u/1318618870"

	bytes, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./citylist_test_data.html", bytes, 0777)

	if err != nil {
		panic(err)
	}
}

func TestParserCityList(t *testing.T) {
	contents, err := os.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	cityList := parser.ParseCityList(contents)

	if len(cityList.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(cityList.Requests))
	}
	for i, url := range expectedUrls {
		if url != cityList.Requests[i].Url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, cityList.Requests[i].Url)
		}
	}

	if len(cityList.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(cityList.Items))
	}
	for i, item := range expectedCities {
		if item != cityList.Items[i] {
			t.Errorf("expected citylist #%d: %s, but was %s", i, item, cityList.Items[i])
		}
	}
}
