package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/model"
	"testing"
)

func Test_save(t *testing.T) {
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


	// TODO:try to start up elasic search
	// here using docker go client
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	index := "dating_profile_test"
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got:%v, expected:%v",actual, expected)
	}
}
