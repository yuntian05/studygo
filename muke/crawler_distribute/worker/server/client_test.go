package main

import (
	"fmt"
	"studygo/muke/crawler_distribute/config"
	"studygo/muke/crawler_distribute/rpcsupport"
	"studygo/muke/crawler_distribute/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9001"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second * 2)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	
	req := worker.Request{
		Url:    "https://album.zhenai.com/u/1657582882",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "一杯冰美式",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRPC, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}