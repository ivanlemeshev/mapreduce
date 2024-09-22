package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/ivanlemeshev/mapreduce/internal/mapreduce/distributed"
	apiv1 "github.com/ivanlemeshev/mapreduce/pkg/maprecude/api/v1"
)

func main() {
	log.Println("Starting the worker server...")

	l, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	workerServer := distributed.NewWorkerServer()

	apiv1.RegisterWorkerServiceServer(grpcServer, workerServer)

	log.Println("Worker server is running on 127.0.0.1:8001...")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
