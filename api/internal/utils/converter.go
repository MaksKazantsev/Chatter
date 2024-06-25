package utils

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
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

func (c converter) CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq {
	return &messagesPkg.CreateMessageReq{Token: req.Token, SenderID: req.SenderID, ReceiverID: req.ReceiverID, ChatID: req.ChatID, Value: req.Value, ReceiverOffline: receiverOffline}
}

func (c converter) DeleteMsgToPb(messageID string) *messagesPkg.DeleteMessageReq {
	return &messagesPkg.DeleteMessageReq{MessageID: messageID}
}

func (c converter) UpdateTokensToPb(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{RToken: req}
}
func (c converter) ParseTokenToPb(req string) *userPkg.ParseTokenReq {
	return &userPkg.ParseTokenReq{Token: req}
}

func (c converter) SuggestFs(req models.FriendShipReq) *userPkg.SuggestFriendShipReq {
	return &userPkg.SuggestFriendShipReq{Token: req.Token, Receiver: req.Receiver}
}

func (c converter) RefuseFs(req models.RefuseFriendShipReq) *userPkg.RefuseFriendShipReq {
	return &userPkg.RefuseFriendShipReq{Token: req.Token, Sender: req.Sender}
}

func (c converter) UpdateTokens(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{
		RToken: req,
	}
}

type ToPb interface {
	RegReqToPb(req models.SignupReq) *userPkg.RegisterReq
	LogReqToPb(req models.LoginReq) *userPkg.LoginReq
	SendCodeReqToPb(req string) *userPkg.SendReq
	VerifyCodeReqToPb(req models.VerifyCodeReq) *userPkg.VerifyReq
	RecoveryReqToPb(req models.RecoveryReq) *userPkg.RecoveryReq
	UpdateTokensToPb(req string) *userPkg.UpdateTokenReq
	ParseTokenToPb(req string) *userPkg.ParseTokenReq
	DeleteMsgToPb(messageID string) *messagesPkg.DeleteMessageReq
	CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq
}

type ToService interface {
}

func (c converter) RecoveryReqToPb(req models.RecoveryReq) *userPkg.RecoveryReq {
	return &userPkg.RecoveryReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c converter) VerifyCodeReqToPb(req models.VerifyCodeReq) *userPkg.VerifyReq {
	return &userPkg.VerifyReq{
		Code:  req.Code,
		Email: req.Email,
		Type:  req.Type,
	}
}

func (c converter) SendCodeReqToPb(email string) *userPkg.SendReq {
	return &userPkg.SendReq{
		Email: email,
	}
}

func (c converter) LogReqToPb(req models.LoginReq) *userPkg.LoginReq {
	return &userPkg.LoginReq{Email: req.Email, Password: req.Password}
}
func (c converter) RegReqToPb(req models.SignupReq) *userPkg.RegisterReq {
	return &userPkg.RegisterReq{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}
