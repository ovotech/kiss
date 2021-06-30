package aws

import (
	"errors"
	"fmt"

	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	smtypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go"
	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	"github.com/ovotech/kiss/pkg/ref"
)

// TODO modify this to support multiple bindings to the same secret. Right now we can only support
// one secret.
func (m *Manager) makeSecretPolicy(serviceAccountARN string) string {
	return fmt.Sprintf(`{
  "Version" : "2012-10-17",
  "Statement" : [
    {
      "Effect": "Allow",
      "Principal": {"AWS": "%s"},
      "Action": "secretsmanager:GetSecretValue",
      "Resource": "*",
      "Condition": {
        "ForAnyValue:StringEquals": {
          "secretsmanager:VersionStage" : "AWSCURRENT"
        }
      }
    }
  ]
}`, serviceAccountARN)
}

// Create a secret with the given string value. The secret will be logically scoped to the provided
// namespace (i.e. will only bind to service account roles also in that namespace).
func (m *Manager) CreateSecret(namespace, name, value string) error {
	secretName := m.makeSecretName(namespace, name)
	tags := []smtypes.Tag{
		{Key: ref.String(managedByTagKey), Value: ref.String(managedByTagValue)},
		{Key: ref.String(namespaceTagKey), Value: &namespace},
		{Key: ref.String(nameTagKey), Value: &name},
	}

	_, err := m.smclient.CreateSecret(
		m.ctx,
		&sm.CreateSecretInput{Name: &secretName, SecretString: &value, Tags: tags},
	)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceExistsException" {
			return &awserrors.AWSError{
				Code:    awserrors.AlreadyExistsErrorCode,
				Message: "A resource with this name already exists.",
			}
		}
		if errors.As(err, &ae) && ae.ErrorCode() == "InvalidRequestException" {
			return &awserrors.AWSError{
				Code:    awserrors.InvalidRequestErrorCode,
				Message: ae.ErrorMessage(),
			}
		}
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return nil
}

// Gets the secret with namesapce and name.
func (m *Manager) GetSecret(namespace, name string) (*sm.DescribeSecretOutput, error) {
	secretName := m.makeSecretName(namespace, name)

	secretOutput, err := m.smclient.DescribeSecret(
		m.ctx,
		&sm.DescribeSecretInput{SecretId: &secretName},
	)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			return nil, &awserrors.AWSError{
				Code:    awserrors.NotFoundErrorCode,
				Message: "Couldn't find a secret with this name.",
			}
		}
		return nil, &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return secretOutput, nil
}

func (m *Manager) isManagedSecret(secretOutput *sm.DescribeSecretOutput) bool {
	for _, tag := range secretOutput.Tags {
		if *tag.Key == managedByTagKey && *tag.Value == managedByTagValue {
			return true
		}
	}

	return false
}

// Delete a secret with the given name.
func (m *Manager) DeleteSecret(namespace, name string) error {
	secretName := m.makeSecretName(namespace, name)

	secret, err := m.GetSecret(
		namespace,
		name,
	)
	if err != nil {
		return err
	}

	if !m.isManagedSecret(secret) {
		return &awserrors.AWSError{
			Code:    awserrors.NotManagedErrorCode,
			Message: "The secret is not managed by KISS",
		}
	}

	_, err = m.smclient.DeleteSecret(
		m.ctx,
		&sm.DeleteSecretInput{SecretId: &secretName, ForceDeleteWithoutRecovery: true},
	)

	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "ResourceNotFoundException" {
			return &awserrors.AWSError{
				Code:    awserrors.NotFoundErrorCode,
				Message: "Couldn't find a secret with this name.",
			}
		}
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}
	return nil
}
