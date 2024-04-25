package errors

type ProductError struct {
	ServerInternalError  error
	ProductNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type ProductNotFoundError struct{}

func (e *ProductNotFoundError) Error() string {
	return "Product not found"
}
