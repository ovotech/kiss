package aws

import "fmt"

// Returns the name for the role, as used in AWS. This is a string with the format:
// (prefix_)namespace_name
func (m *Manager) makeIAMRoleName(namespace, name string) string {
	if m.rolePrefix == "" {
		return fmt.Sprintf("%s_%s", namespace, name)
	}
	return fmt.Sprintf("%s_%s_%s", m.rolePrefix, namespace, name)
}

// makeRoleARN returns the AWS ARN for a role given the k8s ServieAccount namespace/name. Note that
// this is an ARN generated locally from the name and namespace strings and is not an ARN looked up
// on AWS. As such this role may or may not exist in AWS.
func (m *Manager) makeRoleARN(namespace, name string) string {
	roleName := m.makeIAMRoleName(namespace, name)
	return fmt.Sprintf("arn:aws:iam::%s:role/%s", m.accountId, roleName)
}
