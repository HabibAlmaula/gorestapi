package exception

type UnAuthorizedError struct {
	Error string
}

func NewUnAuthorizedError(error string) *UnAuthorizedError {
	return &UnAuthorizedError{Error: error}
}
