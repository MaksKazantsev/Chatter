package utils

import (
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	ValidateCreateMessageReq(req *pkg.CreateMessageReq) error
}

func NewValidator() Validator {
	return &validator{}
}

type validator struct {
}

func (v validator) ValidateCreateMessageReq(req *pkg.CreateMessageReq) error {
	if req.Value == "" || req.Value == " " || len(req.Value) > 100 {
		return status.Error(codes.InvalidArgument, "Invalid message value")
	}
	return nil
}
