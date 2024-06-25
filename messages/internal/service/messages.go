package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	"github.com/google/uuid"
	"time"
)

type Messages struct {
	repo db.Messages
}

func NewMessages(repo db.Messages) *Messages {
	return &Messages{repo: repo}
}

func (m *Messages) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {

	msg.MessageID = uuid.New().String()
	msg.SentAt = time.Now()

	log.GetLogger(ctx).Info("Service layer success")

	if err := m.repo.CreateMessage(ctx, msg, receiverOffline); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
func (m *Messages) DeleteMessage(ctx context.Context, messageID string) error {
	log.GetLogger(ctx).Info("Service layer success")
	if err := m.repo.DeleteMessage(ctx, messageID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (m *Messages) GetHistory(ctx context.Context, req models.GetHistoryReq, uuid string) ([]models.Message, error) {
	log.GetLogger(ctx).Info("Service layer success")
	res, err := m.repo.GetHistory(ctx, req, uuid)
	if err != nil {
		return nil, fmt.Errorf("repo error: %w", err)
	}
	return res, nil
}
