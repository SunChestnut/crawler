package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatalln(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", config.WorkerPort0),
			worker.CrawlService{},
		),
	)
}
