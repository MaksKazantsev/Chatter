package utils

import (
	"github.com/MaksKazantsev/SSO/api/internal/models"
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
)

type Converter interface {
	ToPb
	ToService
}

func NewConverter() Converter {
	return converter{}
}

type converter struct {
}

type ToPb interface {
	RegReqToPb(req models.SignupReq) *pkg.RegisterReq
	LogReqToPb(req models.LoginReq) *pkg.LoginReq
	ResReqToPb(req models.ResetReq) *pkg.ResetReq
}

type ToService interface {
}

func (c converter) ResReqToPb(req models.ResetReq) *pkg.ResetReq {
	return &pkg.ResetReq{OldPassword: req.OldPassword, NewPassword: req.NewPassword}
}
func (c converter) LogReqToPb(req models.LoginReq) *pkg.LoginReq {
	return &pkg.LoginReq{Email: req.Email, Password: req.Password}
}
func (c converter) RegReqToPb(req models.SignupReq) *pkg.RegisterReq {
	return &pkg.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}
