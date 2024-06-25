package utils

import (
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type Converter interface {
	ToPb
}

func NewConverter() Converter {
	return &converter{}
}

type converter struct{}

type ToPb interface {
	RegReqToPb(req models.SignupReq) *userPkg.RegisterReq
	LogReqToPb(req models.LoginReq) *userPkg.LoginReq
	SendCodeReqToPb(req string) *userPkg.SendReq
	VerifyCodeReqToPb(req models.VerifyCodeReq) *userPkg.VerifyReq
	RecoveryReqToPb(req models.RecoveryReq) *userPkg.RecoveryReq
	UpdateTokensToPb(req string) *userPkg.UpdateTokenReq
	ParseTokenToPb(req string) *userPkg.ParseTokenReq
	DeleteMsgToPb(messageID, token string) *messagesPkg.DeleteMessageReq
	CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq
}

// User Microservice

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

func (c converter) UpdateTokensToPb(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{RToken: req}
}
func (c converter) ParseTokenToPb(req string) *userPkg.ParseTokenReq {
	return &userPkg.ParseTokenReq{Token: req}
}
func (c converter) UpdateTokens(req string) *userPkg.UpdateTokenReq {
	return &userPkg.UpdateTokenReq{
		RToken: req,
	}
}

// Messages Microservice
/*
func (c converter) MessageToService(messages []*messagesPkg.Message) []models.Message {
	var msg models.Message
	msgs := make([]models.Message, len(messages))
	for i, _ := range messages {
		msg.MessageID = messages[i].MessageID
		msg.SenderID = messages[i].SenderID
		msg.ReceiverID = messages[i].ReceiverID
		msg.Value = messages[i].Value
		msg.ChatID = messages[i].ChatID
		msg.SentAt = messages[i].SentAt.AsTime()
		msgs = append(msgs, msg)
	}
	return msgs
}

*/

func (c converter) CreateMsgToPb(req *models.Message, receiverOffline bool) *messagesPkg.CreateMessageReq {
	return &messagesPkg.CreateMessageReq{Token: req.Token, SenderID: req.SenderID, ReceiverID: req.ReceiverID, ChatID: req.ChatID, Value: req.Value, ReceiverOffline: receiverOffline}
}

func (c converter) DeleteMsgToPb(messageID, token string) *messagesPkg.DeleteMessageReq {
	return &messagesPkg.DeleteMessageReq{MessageID: messageID, Token: token}
}
