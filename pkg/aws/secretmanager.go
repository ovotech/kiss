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

// Bind secret to a given service account. This modifies the secret's resource-based policy to allow
// it to be read by the relevant service account role.
func (m *Manager) BindSecret(namespace, name, serviceAccountName string) error {
	secret, err := m.GetSecret(namespace, name)
	if err != nil {
		return err
	}

	if !m.isManagedSecret(secret) {
		return &awserrors.AWSError{
			Code:    awserrors.NotManagedErrorCode,
			Message: "The resource is not managed by this service. Refusing to modify it.",
		}
	}

	secretName := secret.Name
	serviceAccountARN := m.makeRoleARN(namespace, serviceAccountName)
	policy := m.makeSecretPolicy(serviceAccountARN)

	_, err = m.smclient.PutResourcePolicy(
		m.ctx,
		&sm.PutResourcePolicyInput{ResourcePolicy: &policy, SecretId: secretName},
	)
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "MalformedPolicyDocumentException" {
			return &awserrors.AWSError{
				Code:    awserrors.MalformedPolicyErrorCode,
				Message: "Got a malformed policy error from AWS. This can happen when the service account role doesn't exist.",
			}
		}
		return &awserrors.AWSError{Code: awserrors.OtherErrorCode, Message: err.Error()}
	}

	return nil
}

func (m *Manager) isManagedSecret(secretOutput *sm.DescribeSecretOutput) bool {
	for _, tag := range secretOutput.Tags {
		if *tag.Key == managedByTagKey && *tag.Value == managedByTagValue {
			return true
		}
	}

	return false
}
