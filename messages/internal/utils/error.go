package utils

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func HandleError(err error) error {
	var e *Error
	if !errors.As(err, &e) {
		return status.Error(codes.Internal, err.Error())
	}
	switch e.Status {
	case ErrNotFound:
		return status.Error(codes.NotFound, e.Message)
	case ErrNotAllowed:
		return status.Error(codes.PermissionDenied, e.Message)
	case ErrBadRequest:
		return status.Error(codes.InvalidArgument, e.Message)
	case ErrInternal:
		return status.Error(codes.Internal, err.Error())

	default:
		return status.Error(codes.Internal, fmt.Sprintf("Unexpected: %s", err.Error()))
	}
}
