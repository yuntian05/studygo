package main

import (
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/model"
	"studygo/muke/crawler_distribute/config"
	"studygo/muke/crawler_distribute/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T)  {
	const host = ":1234"
	// start ItemSaverServer
	//go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	expected := engine.Item{
		Id:"1657582882",
		Url:"https://album.zhenai.com/u/1657582882",
		Type:"zhenai",
		Payload: model.Profile{
			Name:       "一杯冰美式",
			Marriage:   "未婚",
			Age:        "34岁",
			Xingzuo:    "魔羯座(12.22-01.19)",
			Height:     "168cm",
			Weight:     "55kg",
			Workplace:  "工作地:北京顺义区",
			Income:     "月收入:2-5万",
			Occupation: "其他职业",
			Education:  "大学本科",
		},
	}

	result := ""
	err = client.Call(config.ItemSaveRPC, expected, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%v error:%v", result, err)
	}
}