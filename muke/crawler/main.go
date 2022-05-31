package main

import (
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/persist"
	"studygo/muke/crawler/scheduler"
	"studygo/muke/crawler/zhenai/parser"
	"studygo/muke/crawler_distribute/config"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/beijing",
	//	Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	//})
}
