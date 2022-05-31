package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"studygo/muke/crawler/engine"
	"studygo/muke/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error  {
	err := persist.Save(s.Client, s.Index, item)
	if err != nil {
		log.Printf("Error:%v %v", item, err)
		return err
	}
	log.Printf("item %v saved", item)
	*result = "ok"
	return nil
}