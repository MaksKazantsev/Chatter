package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
)

func (p *Postgres) CreateMessage(ctx context.Context, msg *models.Message) error {
	q := `INSERT INTO messages(chatid,senderid,receiverid,messageid,val,sentat)`
	_, err := p.Exec(q, msg.ChatID, msg.SenderID, msg.ReceiverID, msg.MessageID, msg.Value, msg.SentAt)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}

func (p *Postgres) DeleteMessage(ctx context.Context, messageID string) error {
	q := `DELETE FROM messages WHERE messageid = $1`
	_, err := p.Exec(q, messageID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("Message does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}
