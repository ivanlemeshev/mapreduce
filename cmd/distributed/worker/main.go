package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	apiv1 "github.com/ivanlemeshev/mapreduce/pkg/maprecude/api/v1"
)

func main() {
	log.Println("Starting the worker server...")

	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	workerServer := NewWorkerServer()

	apiv1.RegisterWorkerServiceServer(grpcServer, workerServer)

	log.Println("Worker server is running on port 8001...")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

type WorkerServer struct {
	apiv1.UnimplementedWorkerServiceServer
}

func NewWorkerServer() *WorkerServer {
	return &WorkerServer{}
}

func (s *WorkerServer) ServerLive(
	ctx context.Context,
	req *apiv1.ServerLiveRequest,
) (*apiv1.ServerLiveResponse, error) {
	return &apiv1.ServerLiveResponse{
		IsLive: true,
	}, nil
}
