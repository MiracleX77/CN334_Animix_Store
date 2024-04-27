package errors

type ReviewError struct {
	ServerInternalError error
	ReviewNotFoundError string
}

type ServerInternalError struct {
	Err error
}

func (e *ServerInternalError) Error() string {
	return "server internal error: " + e.Err.Error()
}

type ReviewNotFoundError struct{}

func (e *ReviewNotFoundError) Error() string {
	return "Review not found"
}
