package persist

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestItemSaver(t *testing.T) {

	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.56.10:9200/"},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		t.Errorf("Connection Failure : %v", err)
	}
	const index = "dating_profile"

	item := engine.Item{
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
	if err := save(client, item, index); err != nil {
		t.Errorf("An error occured %v", err)
	}

	// TODO 如何根据 DocumentID 检索数据？
	// 若根据 ID 从 ElasticSearch 检索的数据与上述数据相同，则算保存成功
	actual := search(client, index, item.Id)
	if item != actual {
		t.Errorf("🥰Expected %v; 😭Got %v", item, actual)
	}
}
