package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)

// 需在命令行中启动下述的 Main 函数，且加上端口参数：go run worker.go -port 8899
var port = flag.Int("port", 0, "the port for me to listen on")

// go run worker.go -port 9000
// go run worker.go -port 9001
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("💔must specify a port")
		return
	}
	log.Fatalln(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port),
			worker.CrawlService{},
		),
	)
}
