package test

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"crawler/concurrent/zhenai/parser"
	"os"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := os.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parser.ParseProfile(contents, "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481", "一身傲气如你*")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481",
		Id:  "3903982005871861481",
		Payload: model.Profile{
			Name:       "一身傲气如你*",
			Age:        24,
			Height:     140,
			Weight:     257,
			Income:     "8001-10000元",
			Gender:     "男",
			XinZuo:     "天蝎座",
			Occupation: "测试工程师",
			Marriage:   "未婚",
			House:      "无房",
			HuKou:      "沈阳市",
			Education:  "硕士",
			Car:        "有车",
		},
	}

	if expected != actual {
		t.Errorf("expected %v; but was %v", expected, actual)
	}

}
