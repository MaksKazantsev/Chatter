package utils

import (
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type Converter interface {
	ToPb
	ToService
}

type ToPb interface {
	ParseTokenToPb(token string) *pkg.ParseTokenReq
}
type ToService interface {
}

type converter struct {
}

func (c converter) ParseTokenToPb(token string) *pkg.ParseTokenReq {
	return &pkg.ParseTokenReq{Token: token}
}

func NewConverter() Converter {
	return &converter{}
}
