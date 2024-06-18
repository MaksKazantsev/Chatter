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
}

type ToService interface {
	RegReqToService(req *pkg.RegisterReq) models.RegReq
	LoginReqToService(req *pkg.LoginReq) models.LogReq
	ResetReqToService(req *pkg.ResetReq) models.ResReq
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct {
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

func (c converter) ResetReqToService(req *pkg.ResetReq) models.ResReq {
	return models.ResReq{OldPassword: req.OldPassword, Password: req.NewPassword, Token: req.Token}
}
