package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles UpdateSecretRequests
func (s *kissServer) UpdateSecret(
	ctx context.Context,
	updateSecretRequest *pb.UpdateSecretRequest,
) (*pb.UpdateSecretResponse, error) {
	if !isValidNameAndNamespace(updateSecretRequest.Metadata.Namespace, updateSecretRequest.Name) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			updateSecretRequest.Metadata.Namespace,
			updateSecretRequest.Name,
		)
	}

	err := AWSManager.UpdateSecret(
		updateSecretRequest.Metadata.Namespace,
		updateSecretRequest.Name,
		updateSecretRequest.Value,
	)

	if err != nil {
		log.Info().Msgf("Error updating secret: %v", err)
		if awserrors.IsNotFound(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"resource '%s/%s' not found",
				updateSecretRequest.Metadata.Namespace,
				updateSecretRequest.Name,
			)
		}
		if awserrors.IsInvalidRequest(err) {
			return nil, status.Errorf(
				codes.FailedPrecondition,
				"resource '%s/%s' could not be updated: %s",
				updateSecretRequest.Metadata.Namespace,
				updateSecretRequest.Name,
				err,
			)
		}
		return nil, status.Errorf(codes.Unknown, "failed to update secret for unknown reasons")
	}

	log.Info().
		Msgf("Updated secret '%s/%s' on behalf of '%s'", updateSecretRequest.Metadata.Namespace, updateSecretRequest.Name, ctx.Value("user"))
	return &pb.UpdateSecretResponse{}, nil
}
