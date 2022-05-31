package parser

import (
	"regexp"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler_distribute/config"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matchs := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		//result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	matchs = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}
	return result
}
