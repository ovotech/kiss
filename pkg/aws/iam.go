package aws

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamtypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go"
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
			"Resource": "%s",
			"Condition": {
				"ForAnyValue:StringEquals": {
					"secretsmanager:VersionStage" : "AWSCURRENT"
				}
      		}
		}
	}`, arn)
}

// Create an IAM policy that allows reading a secret with the provided namespace/name and ARN
func (m *Manager) CreateSecretIAMPolicy(namespace, name, arn string) error {
	policyName := m.makeSecretPolicyName(namespace, name)
	policy := m.makeSecretIAMPolicy(arn)

	tags := []iamtypes.Tag{
		{Key: ref.String(managedByTagKey), Value: ref.String(managedByTagValue)},
		{Key: ref.String(namespaceTagKey), Value: &namespace},
		{Key: ref.String(nameTagKey), Value: &name},
	}

	_, err := m.iamclient.CreatePolicy(
		m.ctx,
		&iam.CreatePolicyInput{PolicyDocument: &policy, PolicyName: &policyName, Tags: tags},
	)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "EntityAlreadyExists" {
			return &awserrors.AWSError{
				Code:    awserrors.AlreadyExistsErrorCode,
				Message: "An IAM policy for this secret already exists",
			}
		}
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return nil
}

// Attach a secret's IAM policy to a service account IAM role
func (m *Manager) AttachIAMPolicy(namespace, name, serviceAccountName string) error {
	policyARN := m.makeSecretPolicyARN(namespace, name)
	serviceAccountRoleName := m.makeServiceAccountRoleName(namespace, serviceAccountName)

	_, err := m.iamclient.AttachRolePolicy(
		m.ctx,
		&iam.AttachRolePolicyInput{PolicyArn: &policyARN, RoleName: &serviceAccountRoleName},
	)
	if err != nil {
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return nil
}

// Gets the IAM policy with namesapce and name.
func (m *Manager) GetIamPolicy(namespace, name string) (*iam.GetPolicyOutput, error) {
	policyARN := m.makeSecretPolicyARN(namespace, name)

	policyOutput, err := m.iamclient.GetPolicy(
		m.ctx,
		&iam.GetPolicyInput{PolicyArn: &policyARN},
	)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			return nil, &awserrors.AWSError{
				Code:    awserrors.NotFoundErrorCode,
				Message: "Couldn't find a IAM policy for this secret.",
			}
		}
		return nil, &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return policyOutput, nil
}

func (m *Manager) isManagedPolicy(policyOutput *iam.GetPolicyOutput) bool {
	for _, tag := range policyOutput.Policy.Tags {
		if *tag.Key == managedByTagKey && *tag.Value == managedByTagValue {
			return true
		}
	}

	return false
}

// Delete IAM Policy for a secret with the given name.
func (m *Manager) DeleteSecretIAMPolicy(namespace, name string) error {
	policy, err := m.GetIamPolicy(
		namespace,
		name,
	)
	if err != nil {
		return err
	}

	if !m.isManagedPolicy(policy) {
		return &awserrors.AWSError{
			Code:    awserrors.NotManagedErrorCode,
			Message: "The policy is not managed by KISS",
		}
	}

	_, err = m.iamclient.DeletePolicy(
		m.ctx,
		&iam.DeletePolicyInput{PolicyArn: policy.Policy.Arn},
	)

	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			return &awserrors.AWSError{
				Code:    awserrors.NotFoundErrorCode,
				Message: "Couldn't find a IAM policy for this secret.",
			}
		}
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}
	return nil
}
