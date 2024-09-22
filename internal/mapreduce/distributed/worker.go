package distributed

import (
	"context"

	apiv1 "github.com/ivanlemeshev/mapreduce/pkg/maprecude/api/v1"
)

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
