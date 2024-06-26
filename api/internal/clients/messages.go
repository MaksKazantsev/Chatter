package clients

import (
	"context"
	"github.com/MaksKazantsev/Chatter/api/internal/models"
	"github.com/MaksKazantsev/Chatter/api/internal/utils"
	pkg "github.com/MaksKazantsev/Chatter/messages/pkg/grpc"
)

type Messages interface {
	CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error
	DeleteMessage(ctx context.Context, messageID, token string) error
	GetHistory(ctx context.Context, req models.GetHistoryReq) ([]models.Message, error)
}

func NewMessages(cl pkg.MessagesClient) Messages {
	return &messagesCl{c: utils.NewConverter(), cl: cl}
}

type messagesCl struct {
	c  utils.Converter
	cl pkg.MessagesClient
}

func (m *messagesCl) DeleteMessage(ctx context.Context, messageID, token string) error {
	if _, err := m.cl.DeleteMessage(ctx, m.c.DeleteMsgToPb(messageID, token)); err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (m *messagesCl) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {
	if _, err := m.cl.CreateMessage(ctx, m.c.CreateMsgToPb(msg, receiverOffline)); err != nil {
		return utils.GRPCErrorToError(err)
	}
	return nil
}

func (m *messagesCl) GetHistory(ctx context.Context, req models.GetHistoryReq) ([]models.Message, error) {
	res, err := m.cl.GetHistory(ctx, m.c.GetHistoryToPb(req))
	if err != nil {
		return nil, utils.GRPCErrorToError(err)
	}
	return m.c.MessageToService(res.Messages), nil
}
