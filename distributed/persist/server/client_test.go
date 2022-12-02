package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	// 1⃣️Start ItemSaverServer
	go serverRpc(host, "test1")
	time.Sleep(200 * time.Millisecond)

	// 2⃣️Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Errorf("error create rpc client : %v", err)
	}

	// 3⃣️Call save method
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
	var result string
	err = client.Call(config.ItemSaverServiceRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("error call ItemSaverService : %v", err)
	}
}
