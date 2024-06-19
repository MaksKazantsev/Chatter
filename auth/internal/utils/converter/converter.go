package converter

import (
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	pkg "github.com/MaksKazantsev/SSO/auth/pkg/grpc"
)

type Converter interface {
	ToPb
	ToService
}

type ToPb interface {
	RegResToPb(req models.RegRes) *pkg.RegisterRes
	LoginResToPb(access, refresh string) *pkg.LoginRes
	VerifyCodeResToPb(access, refresh string) *pkg.VerifyRes
	UpdateTokensResToPb(access, refresh string) *pkg.UpdateTokenRes
}

type ToService interface {
	RegReqToService(req *pkg.RegisterReq) models.RegReq
	LoginReqToService(req *pkg.LoginReq) models.LogReq
	SendCodeReqToService(req *pkg.SendReq) string
	VerifyCodeReqToService(req *pkg.VerifyReq) (string, string, string)
	RecoveryReqToService(req *pkg.RecoveryReq) models.Credentials
	UpdateTokensReqToService(req *pkg.UpdateTokenReq) string
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct {
}

func (c converter) UpdateTokensResToPb(access, refresh string) *pkg.UpdateTokenRes {
	return &pkg.UpdateTokenRes{AToken: access, RToken: refresh}
}

func (c converter) UpdateTokensReqToService(req *pkg.UpdateTokenReq) string {
	return req.RToken
}

func (c converter) RecoveryReqToService(req *pkg.RecoveryReq) models.Credentials {
	return models.Credentials{
		Password: req.Password,
		Email:    req.Email,
	}
}

func (c converter) VerifyCodeResToPb(access, refresh string) *pkg.VerifyRes {
	return &pkg.VerifyRes{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func (c converter) VerifyCodeReqToService(req *pkg.VerifyReq) (string, string, string) {
	return req.Code, req.Email, req.Type
}

func (c converter) SendCodeReqToService(req *pkg.SendReq) string {
	return req.Email
}

func (c converter) RegResToPb(req models.RegRes) *pkg.RegisterRes {
	return &pkg.RegisterRes{UUID: req.UUID, AccessToken: req.AccessToken, RefreshToken: req.RefreshToken}
}

func (c converter) LoginResToPb(access, refresh string) *pkg.LoginRes {
	return &pkg.LoginRes{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func (c converter) RegReqToService(req *pkg.RegisterReq) models.RegReq {
	return models.RegReq{Password: req.Password, Email: req.Email, Username: req.Username}
}

func (c converter) LoginReqToService(req *pkg.LoginReq) models.LogReq {
	return models.LogReq{Email: req.Email, Password: req.Password}
}
