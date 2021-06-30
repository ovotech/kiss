package errors

import "fmt"

const (
	NotFoundErrorCode        = "NotFound"
	NotManagedErrorCode      = "NotManaged"
	OtherErrorCode           = "OtherError"
	AlreadyExistsErrorCode   = "AlreadyExists"
	MalformedPolicyErrorCode = "MalformedPolicy"
	InvalidRequestErrorCode  = "InvalidRequest"
)

type AWSError struct {
	Code    string
	Message string
}

func (e *AWSError) Error() string {
	return fmt.Sprintf("AWSError %s: %s", e.Code, e.Message)
}

// Returns true if the error is a NotFound
func IsNotFound(err error) bool {
	if err, ok := err.(*AWSError); ok && err.Code == NotFoundErrorCode {
		return true
	}
	return false
}

// Returns true if the error is an AlreadyExists
func IsAlreadyExists(err error) bool {
	if err, ok := err.(*AWSError); ok && err.Code == AlreadyExistsErrorCode {
		return true
	}
	return false
}

// Returns true if the error is an MalformedPolicy
func IsMalformedPolicy(err error) bool {
	if err, ok := err.(*AWSError); ok && err.Code == MalformedPolicyErrorCode {
		return true
	}
	return false
}

func IsNotManaged(err error) bool {
	if err, ok := err.(*AWSError); ok && err.Code == NotManagedErrorCode {
		return true
	}
	return false
}

func IsInvalidRequest(err error) bool {
	if err, ok := err.(*AWSError); ok && err.Code == InvalidRequestErrorCode {
		return true
	}
	return false
}
