package exception

type DataExistError struct {
	Error string
}

func NewDataExistError(error string) DataExistError {
	return DataExistError{Error: error}
}
