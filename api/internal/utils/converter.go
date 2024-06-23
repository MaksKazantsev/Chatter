package utils

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	pkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
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

func (c converter) UpdateTokensToPb(req string) *pkg.UpdateTokenReq {
	return &pkg.UpdateTokenReq{RToken: req}
}

func (c converter) SuggestFsToPb(req models.FriendShipReq) *pkg.SuggestFriendShipReq {
	//TODO implement me
	panic("implement me")
}

func (c converter) RefuseFsToPb(req models.RefuseFriendShipReq) *pkg.RefuseFriendShipReq {
	//TODO implement me
	panic("implement me")
}

func (c converter) ParseTokenToPb(req string) *pkg.ParseTokenReq {
	return &pkg.ParseTokenReq{Token: req}
}

func (c converter) SuggestFs(req models.FriendShipReq) *pkg.SuggestFriendShipReq {
	return &pkg.SuggestFriendShipReq{Token: req.Token, Receiver: req.Receiver}
}

func (c converter) RefuseFs(req models.RefuseFriendShipReq) *pkg.RefuseFriendShipReq {
	return &pkg.RefuseFriendShipReq{Token: req.Token, Sender: req.Sender}
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
	UpdateTokensToPb(req string) *pkg.UpdateTokenReq
	SuggestFsToPb(req models.FriendShipReq) *pkg.SuggestFriendShipReq
	RefuseFsToPb(req models.RefuseFriendShipReq) *pkg.RefuseFriendShipReq
	ParseTokenToPb(req string) *pkg.ParseTokenReq
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
