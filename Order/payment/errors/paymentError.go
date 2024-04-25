package errors

type PaymentError struct {
	ServerInternalError  error
	PaymentNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type PaymentNotFoundError struct{}

func (e *PaymentNotFoundError) Error() string {
	return "Payment not found"
}
