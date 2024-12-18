package exception

type DataNotFoundError struct {
	Message string
}

func (e DataNotFoundError) Error() string {
	return e.Message
}

func NewDataNotFoundError(message string) DataNotFoundError {
	return DataNotFoundError{
		Message: message,
	}
}
