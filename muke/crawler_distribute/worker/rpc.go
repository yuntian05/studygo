package worker

import (
	"log"
	"studygo/muke/crawler/engine"
)

type CrawlerService struct {
}

func (CrawlerService) Process(req Request, result *ParseResult) error {
	log.Printf("处理worker")
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}