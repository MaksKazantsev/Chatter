package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const (
	ERR_CLIENT_INVALID_ARGUMENT = iota + 1
	ERR_CLIENT_NOT_FOUND
	ERR_CLIENT_NOT_ALLOWED
	ERR_INTERNAL
)

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string, code int) error {
	return Error{Message: message, Code: code}
}

func HandleError(err error) (int, string) {
	var e *Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError, err.Error()
	}

	switch e.Code {
	case ERR_CLIENT_NOT_FOUND:
		return http.StatusNotFound, e.Message
	case ERR_INTERNAL:
		return http.StatusInternalServerError, e.Message
	case ERR_CLIENT_INVALID_ARGUMENT:
		return http.StatusBadRequest, e.Message
	default:
		return http.StatusInternalServerError, fmt.Sprintf("unexpected server internal error: %w", err)
	}
}

func GRPCErrorToError(err error) error {
	s, ok := status.FromError(err)
	if ok {
		switch s.Code() {
		case codes.InvalidArgument:
			return NewError(s.Message(), ERR_CLIENT_INVALID_ARGUMENT)
		case codes.NotFound:
			return NewError(s.Message(), ERR_CLIENT_NOT_FOUND)
		case codes.Internal:
			return NewError(s.Message(), ERR_INTERNAL)
		case codes.PermissionDenied:
			return NewError(s.Message(), ERR_CLIENT_NOT_ALLOWED)
		default:
			return NewError(s.Message(), ERR_INTERNAL)
		}
	}
	return NewError("unexpected internal error", ERR_INTERNAL)
}
