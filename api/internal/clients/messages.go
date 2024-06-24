package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
)

type Messages interface {
	CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error
	DeleteMessage(ctx context.Context, messageID string) error
}

func NewMessages() Messages {
	return &messagesCl{}
}

type messagesCl struct {
	c utils.Converter
}

func (m *messagesCl) DeleteMessage(ctx context.Context, messageID string) error {
	//TODO implement me
	panic("implement me")
}

func (m *messagesCl) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {
	//TODO implement me
	panic("implement me")
}
