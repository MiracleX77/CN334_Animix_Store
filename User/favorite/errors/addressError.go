package errors

type AddressError struct {
	ServerInternalError  error
	AddressNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type AddressNotFoundError struct{}

func (e *AddressNotFoundError) Error() string {
	return "Address not found"
}
