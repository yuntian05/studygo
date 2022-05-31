package client

import (
	"log"
	"net/rpc"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler_distribute/config"
	"studygo/muke/crawler_distribute/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		c := <- clientChan
		log.Printf("开始调用crawl rpc")
		err := c.Call(config.CrawlServiceRPC, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeResult(sResult), nil
	}
}
