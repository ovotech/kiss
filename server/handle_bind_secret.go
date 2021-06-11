package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles BindSecretRequests
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

	err := AWSManager.BindSecret(
		bindSecretRequest.Metadata.Namespace,
		bindSecretRequest.Name,
		bindSecretRequest.ServiceAccountName,
	)

	if err != nil {
		log.Info().Msgf("Error creating secret: %v", err)
		if awserrors.IsMalformedPolicy(err) {
			return nil, status.Error(
				codes.InvalidArgument,
				"Got a malformed policy error from AWS. This can happen when the service account role doesn't exist. Check the service account name and contact an admin if the problem perists.",
			)
		} else if awserrors.IsNotFound(err) {
			return nil, status.Error(
				codes.NotFound,
				"Couldn't find a secret with the supplied name.",
			)
		}
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
