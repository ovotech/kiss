package server

import (
	"context"

	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles BindSecretRequests
// Attaches a policy to a ServiceAccount IAM role.
func (s *kissServer) BindSecret(
	ctx context.Context,
	bindSecretRequest *pb.BindSecretRequest,
) (*pb.BindSecretResponse, error) {
	if !isValidNameAndNamespace(bindSecretRequest.Metadata.Namespace, bindSecretRequest.Name) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			bindSecretRequest.Metadata.Namespace,
			bindSecretRequest.Name,
		)
	}

	err := AWSManager.AttachIAMPolicy(
		bindSecretRequest.Metadata.Namespace,
		bindSecretRequest.Name,
		bindSecretRequest.ServiceAccountName,
	)

	if err != nil {
		log.Info().Msgf("Error creating secret: %v", err)
		return nil, status.Errorf(codes.Unknown, "failed to bind secret for unknown reasons")
	}

	log.Info().Msgf(
		"Bound secret '%s/%s' to service account '%s' on behalf of '%s'",
		bindSecretRequest.Metadata.Namespace,
		bindSecretRequest.Name,
		bindSecretRequest.ServiceAccountName,
		ctx.Value("user"),
	)
	return &pb.BindSecretResponse{}, nil
}
