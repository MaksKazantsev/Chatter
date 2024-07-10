package utils

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	messagesPkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
	userPkg "github.com/MaksKazantsev/Chatter/user/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Converter interface {
	ToPb
	ToService
	Internal
}

type ToPb interface {
	ParseTokenToPb(token string) *userPkg.ParseTokenReq
}
type ToService interface {
	CreateMessageToService(req *messagesPkg.CreateMessageReq) (*models.Message, bool)
	GetHistoryToService(req *messagesPkg.GetHistoryReq) models.GetHistoryReq
}

type Internal interface {
	MessageToPb(messages []models.Message) *messagesPkg.GetHistoryRes
}
type converter struct {
}

func (c converter) MessageToPb(messages []models.Message) *messagesPkg.GetHistoryRes {
	var msgs []*messagesPkg.Message
	for i := 0; i < len(messages); i++ {
		msg := &messagesPkg.Message{
			MessageID:  messages[i].MessageID,
			SenderID:   messages[i].SenderID,
			ReceiverID: messages[i].ReceiverID,
			Value:      messages[i].Value,
			ChatID:     messages[i].ChatID,
			SentAt:     timestamppb.New(messages[i].SentAt),
		}
		msgs = append(msgs, msg)
	}

	return &messagesPkg.GetHistoryRes{
		Messages: msgs,
	}
}

func (c converter) GetHistoryToService(req *messagesPkg.GetHistoryReq) models.GetHistoryReq {
	return models.GetHistoryReq{Token: req.Token, ChatID: req.ChatID}
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
