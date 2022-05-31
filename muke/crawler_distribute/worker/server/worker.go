package main

import (
	"flag"
	"fmt"
	"log"
	"studygo/muke/crawler_distribute/rpcsupport"
	"studygo/muke/crawler_distribute/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlerService{}))
}
