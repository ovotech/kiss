package server

import (
	"context"

	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handles CreateSecretIAMPolicyRequests
func (s *kissServer) CreateSecretIAMPolicy(
	ctx context.Context,
	createSecretIAMPolicyRequest *pb.CreateSecretIAMPolicyRequest,
) (*pb.CreateSecretIAMPolicyResponse, error) {
	if !isValidNameAndNamespace(
		createSecretIAMPolicyRequest.Metadata.Namespace,
		createSecretIAMPolicyRequest.Name,
	) {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"invalid input '%s/%s'",
			createSecretIAMPolicyRequest.Metadata.Namespace,
			createSecretIAMPolicyRequest.Name,
		)
	}

	secret, err := AWSManager.GetSecret(
		createSecretIAMPolicyRequest.Metadata.Namespace,
		createSecretIAMPolicyRequest.Name,
	)
	if err != nil {
		log.Info().Msgf("Error creating secret policy: %v", err)
		if awserrors.IsNotFound(err) {
			return nil, status.Errorf(
				codes.NotFound,
				"No secret for '%s/%s', check secret name and namespace",
				createSecretIAMPolicyRequest.Metadata.Namespace,
				createSecretIAMPolicyRequest.Name,
			)
		}
		return nil, status.Errorf(
			codes.Unknown,
			"failed to create secret IAM policy for unknown reasons",
		)
	}

	err = AWSManager.CreateSecretIAMPolicy(
		createSecretIAMPolicyRequest.Metadata.Namespace,
		createSecretIAMPolicyRequest.Name,
		*secret.ARN,
	)
	if err != nil {
		log.Info().Msgf("Error creating secret policy: %v", err)
		if awserrors.IsAlreadyExists(err) {
			return nil, status.Errorf(
				codes.AlreadyExists,
				"resource '%s/%s' already exists",
				createSecretIAMPolicyRequest.Metadata.Namespace,
				createSecretIAMPolicyRequest.Name,
			)
		}
		return nil, status.Errorf(
			codes.Unknown,
			"failed to create secret IAM policy for unknown reasons",
		)
	}

	log.Info().Msgf(
		"Created secret IAM policy '%s/%s' on behalf of '%s'",
		createSecretIAMPolicyRequest.Metadata.Namespace,
		createSecretIAMPolicyRequest.Name,
		ctx.Value("user"),
	)
	return &pb.CreateSecretIAMPolicyResponse{}, nil
}
