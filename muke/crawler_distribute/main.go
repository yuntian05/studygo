package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/scheduler"
	"studygo/muke/crawler/zhenai/parser"
	"studygo/muke/crawler_distribute/config"
	itemsaver "studygo/muke/crawler_distribute/persist/client"
	"studygo/muke/crawler_distribute/rpcsupport"
	worker "studygo/muke/crawler_distribute/worker/client"
)

var (
	itemsaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemsaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/beijing",
	//	Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("Error connecting to %s: %v ", h, err)
			continue
		}
		clients = append(clients, client)
		log.Printf("Connected to %s", h)
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _,client := range clients {
				out <- client
			}
		}
	}()
	return out
}
