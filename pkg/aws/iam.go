package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	"github.com/ovotech/kiss/pkg/ref"
)

// Returns a string for an IAM policy that allows reading the secret. Takes the secret name as
// input.
func (m *Manager) makeSecretIAMPolicy(arn string) string {
	return fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": {
			"Effect": "Allow",
			"Action": "secretsmanager:GetSecretValue",
			"Resource": "%s"
		}
	}`, arn)
}

// Create an IAM policy that allows reading a secret with the provided namespace/name and ARN
func (m *Manager) CreateSecretIAMPolicy(namespace, name, arn string) error {
	secretName := m.makeSecretName(namespace, name)
	policy := m.makeSecretIAMPolicy(arn)

	tags := []iamtypes.Tag{
		{Key: ref.String(managedByTagKey), Value: ref.String(managedByTagValue)},
		{Key: ref.String(namespaceTagKey), Value: &namespace},
		{Key: ref.String(nameTagKey), Value: &name},
	}

	_, err := m.iamclient.CreatePolicy(
		m.ctx,
		&iam.CreatePolicyInput{PolicyDocument: &policy, PolicyName: &secretName, Tags: tags},
	)
	if err != nil {
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return nil
}
