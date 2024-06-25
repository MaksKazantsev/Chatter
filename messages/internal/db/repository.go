package db

import (
	"context"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
)

type Repository interface {
	Messages
}

type Messages interface {
	CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error
	DeleteMessage(ctx context.Context, messageID string) error
	GetHistory(ctx context.Context, req models.GetHistoryReq, uuid string) ([]models.Message, error)
}
