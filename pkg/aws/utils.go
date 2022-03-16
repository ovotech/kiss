package aws

import "fmt"

// Returns the name for the role, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeServiceAccountRoleName(namespace, name string) string {
	if m.rolePrefix == "" {
		return fmt.Sprintf("%s_%s", namespace, name)
	}
	return fmt.Sprintf("%s_%s_%s", m.rolePrefix, namespace, name)
}

// makeRoleARN returns the AWS ARN for a role given the k8s ServieAccount namespace/name. Note that
// this is an ARN generated locally from the name and namespace strings and is not an ARN looked up
// on AWS. As such this role may or may not exist in AWS.
func (m *Manager) makeServiceAccountRoleARN(namespace, name string) string {
	roleName := m.makeServiceAccountRoleName(namespace, name)
	return fmt.Sprintf("arn:aws:iam::%s:role/%s", m.accountId, roleName)
}

// Returns the name for the secret's policy, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeSecretPolicyName(namespace, name string) string {
	if m.secretPrefix == "" {
		return fmt.Sprintf("%s_%s", namespace, name)
	}
	return fmt.Sprintf("%s_%s_%s", m.secretPrefix, namespace, name)
}

// Returns the AWS ARN for the secret's policy
func (m *Manager) makeSecretPolicyARN(namespace, name string) string {
	policyName := m.makeSecretPolicyName(namespace, name)
	return fmt.Sprintf("arn:aws:iam::%s:policy/%s", m.accountId, policyName)
}

// Returns the name for the secret, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeSecretName(namespace, name string) string {
	if m.secretPrefix == "" {
		return fmt.Sprintf("%s-%s", namespace, name)
	}
	return fmt.Sprintf("%s-%s-%s", m.secretPrefix, namespace, name)
}
