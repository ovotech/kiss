package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles CreateSecretRequests
func (s *kissServer) CreateSecret(
	ctx context.Context,
	createSecretRequest *pb.CreateSecretRequest,
) (*pb.CreateSecretResponse, error) {
	if !isValidNameAndNamespace(createSecretRequest.Metadata.Namespace, createSecretRequest.Name) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			createSecretRequest.Metadata.Namespace,
			createSecretRequest.Name,
		)
	}

	err := AWSManager.CreateSecret(
		createSecretRequest.Metadata.Namespace,
		createSecretRequest.Name,
		createSecretRequest.Value,
	)

	if err != nil {
		log.Info().Msgf("Error creating secret: %v", err)
		if awserrors.IsAlreadyExists(err) {
			return nil, status.Errorf(
				codes.AlreadyExists,
				"resource '%s/%s' already exists",
				createSecretRequest.Metadata.Namespace,
				createSecretRequest.Name,
			)
		}
		return nil, status.Errorf(codes.Unknown, "failed to create secret for unknown reasons")
	}

	return &pb.CreateSecretResponse{}, nil
}
