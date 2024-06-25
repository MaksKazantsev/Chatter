package utils

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
)

type Converter interface {
	ToPb
	ToService
}

type ToPb interface {
	ParseTokenToPb(token string) *userPkg.ParseTokenReq
}
type ToService interface {
	CreateMessageToService(req *messagesPkg.CreateMessageReq) (*models.Message, bool)
}

type converter struct {
}

func (c converter) CreateMessageToService(req *messagesPkg.CreateMessageReq) (*models.Message, bool) {
	return &models.Message{SenderID: req.SenderID, ReceiverID: req.ReceiverID, Value: req.Value, ChatID: req.ChatID}, req.ReceiverOffline
}

func (c converter) ParseTokenToPb(token string) *userPkg.ParseTokenReq {
	return &userPkg.ParseTokenReq{Token: token}
}

func NewConverter() Converter {
	return &converter{}
}
