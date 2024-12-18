package exception

type AuthenticationCredentialError struct {
	Message string
}

func (e AuthenticationCredentialError) Error() string {
	return e.Message
}

func NewInvalidCredentialError(message string) AuthenticationCredentialError {
	return AuthenticationCredentialError{
		Message: message,
	}
}
