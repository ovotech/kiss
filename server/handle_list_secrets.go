package server

import (
	"context"

	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles ListSecretRequests
func (s *kissServer) ListSecrets(
	ctx context.Context,
	listSecretsRequest *pb.ListSecretsRequest,
) (*pb.ListSecretsResponse, error) {
	if !isValidString(listSecretsRequest.Metadata.Namespace) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s'",
			listSecretsRequest.Metadata.Namespace,
		)
	}

	secrets, err := AWSManager.ListSecrets(listSecretsRequest.Metadata.Namespace)
	if err != nil {
		log.Info().Msgf("Error listing secrets: %v", err)
		return nil, status.Errorf(codes.Unknown, "failed to list secrets")
	}

	log.Info().Msgf(
		"Listed secrets for '%s' on behalf of '%s'",
		listSecretsRequest.Metadata.Namespace,
		ctx.Value("user"),
	)
	return &pb.ListSecretsResponse{Secrets: secrets}, nil
}
