package errors

type TransactionError struct {
	ServerInternalError      error
	TransactionNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type TransactionNotFoundError struct{}

func (e *TransactionNotFoundError) Error() string {
	return "Transaction not found"
}
