package parser

import (
	"regexp"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler_distribute/config"
)

const citylistRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]*)" [^>]+>([^<]+)</a>`
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(citylistRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		//city := string(m[2])
		//result.Items = append(result.Items, "City " + city)
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
		//fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
	}
	//fmt.Println("match found:", len(matchs))
	return result
}
