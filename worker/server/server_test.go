package main

import (
	"context"
	"crawler/config"
	"crawler/grpcsupport"
	"crawler/pb"
	"crawler/worker/parser"
	"crawler/worker/service"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestWorker(t *testing.T) {
	const address = "127.0.0.1:9100"
	go grpcsupport.NewGrpcWorkerServer(config.Network, address, service.NewCrawlService())

	workerClient := grpcsupport.NewWorkerClient(address)

	request := pb.SerializedRequest{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481",
		Parser: &pb.ParserFunc{
			FunctionName: config.ParseProfile,
			Args:         "一身傲气如你*",
		},
	}

	parserResult, err := workerClient.Process(context.Background(), &request)
	require.NoError(t, err)

	result := parser.DeserializeParserResult(parserResult)

	for _, v := range result.Items {
		log.Printf("Id=%v, profile=%v\n", v.Id, v.Payload)
	}
}
