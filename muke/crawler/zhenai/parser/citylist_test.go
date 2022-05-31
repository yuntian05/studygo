package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	results := ParseCityList(contents, "")
	const resultSize = 470

	if len(results.Requests) != resultSize {
		t.Errorf("result should have %d requests,but had %d", resultSize, len(results.Requests))
	}

	if len(results.Items) != resultSize {
		t.Errorf("result should have %d requests,but had %d", resultSize, len(results.Items))
	}

	//expectedUrls := []string {
	//	"http://www.zhenai.com/zhenghun/aba",
	//	"http://www.zhenai.com/zhenghun/akesu",
	//	"http://www.zhenai.com/zhenghun/alashanmeng",
	//}
	//expectedCitys := []string {
	//	"City 阿坝",
	//	"City 阿克苏",
	//	"City 阿拉善盟",
	//}
	//for i, url := range expectedUrls {
	//	if results.Requests[i].Url != url {
	//		t.Errorf("expected url %d,%s;but was %s", i, url, results.Requests[i].Url)
	//	}
	//}

	//for i, city := range expectedCitys {
	//	if results.Items[i].(string) != city {
	//		t.Errorf("expected url %d,%s;but was %s", i, city, results.Items[i].(string))
	//	}
	//}
	//type args struct {
	//	contents []byte
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want engine.ParseResult
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := ParseCityList(tt.args.contents); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("ParseCityList() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
