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
	return &converter{}
}

type converter struct {
}

func (c converter) UpdateTokens(req string) *pkg.UpdateTokenReq {
	return &pkg.UpdateTokenReq{
		RToken: req,
	}
}

type ToPb interface {
	RegReqToPb(req models.SignupReq) *pkg.RegisterReq
	LogReqToPb(req models.LoginReq) *pkg.LoginReq
	SendCodeReqToPb(req string) *pkg.SendReq
	VerifyCodeReqToPb(req models.VerifyCodeReq) *pkg.VerifyReq
	RecoveryReqToPb(req models.RecoveryReq) *pkg.RecoveryReq
	UpdateTokens(req string) *pkg.UpdateTokenReq
}

type ToService interface {
}

func (c converter) RecoveryReqToPb(req models.RecoveryReq) *pkg.RecoveryReq {
	return &pkg.RecoveryReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c converter) VerifyCodeReqToPb(req models.VerifyCodeReq) *pkg.VerifyReq {
	return &pkg.VerifyReq{
		Code:  req.Code,
		Email: req.Email,
		Type:  req.Type,
	}
}

func (c converter) SendCodeReqToPb(email string) *pkg.SendReq {
	return &pkg.SendReq{
		Email: email,
	}
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
