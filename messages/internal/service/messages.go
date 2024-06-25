package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
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

	if err := m.repo.CreateMessage(ctx, msg, receiverOffline); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
func (m *Messages) DeleteMessage(ctx context.Context, messageID string) error {
	if err := m.repo.DeleteMessage(ctx, messageID); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
