package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"studygo/muke/crawler_distribute/config"
	"studygo/muke/crawler_distribute/persist"
	"studygo/muke/crawler_distribute/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		// must turn off in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
