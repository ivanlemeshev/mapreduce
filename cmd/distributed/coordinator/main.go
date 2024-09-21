package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	apiv1 "github.com/ivanlemeshev/mapreduce/pkg/maprecude/api/v1"
)

func main() {
	log.Println("Starting the coordinator server...")

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	grpcServer := grpc.NewServer()
	coordinatorServer := NewCoordinatorServer()

	apiv1.RegisterCoordinatorServiceServer(grpcServer, coordinatorServer)

	log.Println("Coordinator server is running on port 8000...")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

type CoordinatorServer struct {
	apiv1.UnimplementedCoordinatorServiceServer
}

func NewCoordinatorServer() *CoordinatorServer {
	return &CoordinatorServer{}
}

func (s *CoordinatorServer) ServerLive(
	ctx context.Context,
	req *apiv1.ServerLiveRequest,
) (*apiv1.ServerLiveResponse, error) {
	return &apiv1.ServerLiveResponse{
		IsLive: true,
	}, nil
}
