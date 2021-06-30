package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles DeleteSecretRequests
func (s *kissServer) DeleteSecret(
	ctx context.Context,
	deleteSecretRequest *pb.DeleteSecretRequest,
) (*pb.DeleteSecretResponse, error) {
	if !isValidNameAndNamespace(deleteSecretRequest.Metadata.Namespace, deleteSecretRequest.Name) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			deleteSecretRequest.Metadata.Namespace,
			deleteSecretRequest.Name,
		)
	}

	err := AWSManager.DeleteSecret(
		deleteSecretRequest.Metadata.Namespace,
		deleteSecretRequest.Name,
	)

	if err != nil {
		log.Info().Msgf("Error deleting secret: %v", err)
		if awserrors.IsNotFound(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"resource '%s/%s' not found",
				deleteSecretRequest.Metadata.Namespace,
				deleteSecretRequest.Name,
			)
		}

		if awserrors.IsNotManaged(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"resource '%s/%s' is not managed by KISS",
				deleteSecretRequest.Metadata.Namespace,
				deleteSecretRequest.Name,
			)
		}

		return nil, status.Errorf(codes.Unknown, "failed to delete secret for unknown reasons")
	}

	log.Info().
		Msgf("Deleted secret '%s/%s' on behalf of '%s'", deleteSecretRequest.Metadata.Namespace, deleteSecretRequest.Name, ctx.Value("user"))
	return &pb.DeleteSecretResponse{}, nil
}
