package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles DeleteSecretIAMPolicyRequests
func (s *kissServer) DeleteSecretIAMPolicy(
	ctx context.Context,
	deleteSecretIAMPolicyRequest *pb.DeleteSecretIAMPolicyRequest,
) (*pb.DeleteSecretIAMPolicyResponse, error) {
	if !isValidNameAndNamespace(
		deleteSecretIAMPolicyRequest.Metadata.Namespace,
		deleteSecretIAMPolicyRequest.Name,
	) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			deleteSecretIAMPolicyRequest.Metadata.Namespace,
			deleteSecretIAMPolicyRequest.Name,
		)
	}

	err := AWSManager.DeleteSecretIAMPolicy(
		deleteSecretIAMPolicyRequest.Metadata.Namespace,
		deleteSecretIAMPolicyRequest.Name,
	)
	if err != nil {
		log.Info().Msgf("Error deleting secret IAM Policy: %v", err)
		if awserrors.IsNotFound(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"resource '%s/%s' not found",
				deleteSecretIAMPolicyRequest.Metadata.Namespace,
				deleteSecretIAMPolicyRequest.Name,
			)
		}

		if awserrors.IsNotManaged(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"resource '%s/%s' is not managed by KISS",
				deleteSecretIAMPolicyRequest.Metadata.Namespace,
				deleteSecretIAMPolicyRequest.Name,
			)
		}

		return nil, status.Errorf(codes.Unknown, "failed to delete secret for unknown reasons")
	}

	log.Info().Msgf(
		"Deleted IAM policy for secret '%s/%s' on behalf of '%s'",
		deleteSecretIAMPolicyRequest.Metadata.Namespace,
		deleteSecretIAMPolicyRequest.Name,
		ctx.Value("user"),
	)
	return &pb.DeleteSecretIAMPolicyResponse{}, nil
}
