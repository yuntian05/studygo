package parser

import (
	"io/ioutil"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	results := ParseProfile(contents, "https://album.zhenai.com/u/1657582882", "一杯冰美式")

	actual := results.Items[0]
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

	if expected != actual {
		t.Errorf("got:%v, expected:%v", actual, expected)
	}
}
