package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
)

type Messages interface {
	CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error
	DeleteMessage(ctx context.Context, messageID string) error
}

func NewMessages(cl pkg.MessagesClient) Messages {
	return &messagesCl{c: utils.NewConverter(), cl: cl}
}

type messagesCl struct {
	c  utils.Converter
	cl pkg.MessagesClient
}

func (m *messagesCl) DeleteMessage(ctx context.Context, messageID string) error {
	if _, err := m.cl.DeleteMessage(ctx, m.c.DeleteMsgToPb(messageID)); err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (m *messagesCl) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {
	if _, err := m.cl.CreateMessage(ctx, m.c.CreateMsgToPb(msg)); err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}
