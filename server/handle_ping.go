package server

import (
	"context"

	pb "github.com/ovotech/kiss/proto"
)

// Handler for temporary authorization test method
func (s *kissServer) Ping(
	ctx context.Context,
	pingRequest *pb.PingRequest,
) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}
