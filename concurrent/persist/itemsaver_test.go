package persist

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"testing"
)

func TestItemSaver(t *testing.T) {

	item := engine.Item{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481",
		Id:  "3903982005871861481",
		PayLoad: model.Profile{
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

	if err := save(item); err != nil {
		t.Errorf("An error occured %v", err)
	}

}

func TestSearch(t *testing.T) {
	item := search("3903982005871861481")
	t.Logf("Item : %v", item)
}
