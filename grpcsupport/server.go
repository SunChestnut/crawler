package grpcsupport

import (
	"crawler/pb"
	"crawler/persist/service"
	workerService "crawler/worker/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewGrpcItemSaverServer(network, address string, service *service.ItemSaverService) {
	grpcServer := grpc.NewServer()
	pb.RegisterItemSaverServiceServer(grpcServer, service)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("[grpcsupport.NewGrpcItemSaverServer] cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("[grpcsupport.NewGrpcItemSaverServer] cannot start gprc server: ", err)
	}
}

func NewGrpcWorkerServer(network, address string, service *workerService.CrawlService) {
	grpcServer := grpc.NewServer()
	pb.RegisterCrawlServiceServer(grpcServer, service)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("[grpcsupport.NewGrpcWorkerServer] cannot start server: ", err)
	}
	log.Printf("[grpcsupport.NewGrpcWorkerServer] serer start to listening...\n")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("[grpcsupport.NewGrpcWorkerServer] cannot start gprc server: ", err)
	}
}
