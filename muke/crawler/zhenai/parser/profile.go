package parser

import (
	"log"
	"regexp"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/model"
	"studygo/muke/crawler_distribute/config"
)

const profileReStr = `<div class="m-btn purple"[^>]*>([^<]+)</div>`
const urlReStr = `https://album.zhenai.com/u/([\d]+)`
var profileKey = [10]string{
	"Name",
	"Marriage",
	"Age",
	"Xingzuo",
	"Height",
	"Weight",
	"Worker",
	"Income",
	"Occupation",
	"Education",
}

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	re := regexp.MustCompile(profileReStr)
	matchs := re.FindAllSubmatch(contents, -1)

	profile := &model.Profile{}
	profile.Name = name
	for index, m := range matchs {
		//key := profileKey[index]
		val := string(m[1])
		switch index {
		case 0:
			profile.Marriage = val
		case 1:
			profile.Age = val
		case 2:
			profile.Xingzuo = val
		case 3:
			profile.Height = val
		case 4:
			profile.Weight = val
		case 5:
			profile.Workplace = val
		case 6:
			profile.Income = val
		case 7:
			profile.Occupation = val
		case 8:
			profile.Education = val
		default:
			log.Printf("unknow attribute ")
		}
	}
	urlRe := regexp.MustCompile(urlReStr)
	match := urlRe.FindSubmatch([]byte(url))
	id := ""
	if len(match) >= 2 {
		id = string(match[1])
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id:id,
				Url:url,
				Type:"zhenai",
				Payload: *profile,
			},
		},
	}
	return result
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser  {
	return &ProfileParser{
		userName:name,
	}
}

//func ProfileParser(name string) engine.ParseFunc {
//	return func(c []byte, url string) engine.ParseResult {
//		return ParseProfile(c, url, name)
//	}
//}
