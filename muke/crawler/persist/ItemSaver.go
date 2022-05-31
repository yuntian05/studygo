package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"studygo/muke/crawler/engine"
)

func ItemSaver(index string) (chan engine.Item,error){
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for  {
			item := <- out
			log.Printf("Item Saver: got item: #%d, %v", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Save: err saving item %v:%v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must have type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}