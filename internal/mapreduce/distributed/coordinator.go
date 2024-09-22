package distributed

import (
	"context"

	apiv1 "github.com/ivanlemeshev/mapreduce/pkg/maprecude/api/v1"
)

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
