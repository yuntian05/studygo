package client

import (
	"log"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler_distribute/config"
	"studygo/muke/crawler_distribute/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item,error){
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for  {
			item := <- out
			log.Printf("Item Saver: got item: #%d, %v", itemCount, item)
			// Call RPC to save item
			result := ""
			err = client.Call(config.ItemSaveRPC, item, &result)
			if err != nil || result != "ok"{
				log.Printf("Item Save: err saving item %v:%v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}