package aws

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	smtypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/aws/smithy-go"
	awserrors "github.com/ovotech/kiss/pkg/aws/errors"
	"github.com/ovotech/kiss/pkg/ref"
)

// Returns the name for the role, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeIAMRoleName(namespace, name string) string {
	if m.rolePrefix == "" {
		return fmt.Sprintf("%s_%s", namespace, name)
	}
	return fmt.Sprintf("%s_%s_%s", m.rolePrefix, namespace, name)
}

// Returns the name for the secret, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeSecretName(namespace, name string) string {
	if m.secretPrefix == "" {
		return fmt.Sprintf("%s_%s", namespace, name)
	}
	return fmt.Sprintf("%s_%s_%s", m.secretPrefix, namespace, name)
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

	_, err := m.client.CreateSecret(
		m.ctx,
		&secretsmanager.CreateSecretInput{Name: &secretName, SecretString: &value, Tags: tags},
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
