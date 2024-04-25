package errors

type DeliveryError struct {
	ServerInternalError   error
	DeliveryNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type DeliveryNotFoundError struct{}

func (e *DeliveryNotFoundError) Error() string {
	return "Delivery not found"
}
