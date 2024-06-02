package utils

const (
	ErrBadRequest = iota + 1
	ErrInternal
	ErrNotAllowed
	ErrNotFound
)

type Error struct {
	Message string
	Status  int
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(message string, code int) error {
	return &Error{
		Message: message,
		Status:  code,
	}
}
