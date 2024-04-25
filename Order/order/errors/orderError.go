package errors

type OrderError struct {
	ServerInternalError error
	OrderNotFoundError  string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type OrderNotFoundError struct{}

func (e *OrderNotFoundError) Error() string {
	return "Order not found"
}
