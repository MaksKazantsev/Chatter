package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/Chatter/messages/internal/log"
	"github.com/MaksKazantsev/Chatter/messages/internal/models"
	"github.com/MaksKazantsev/Chatter/messages/internal/utils"
)

func (p *Postgres) CreateMessage(ctx context.Context, msg *models.Message, receiverOffline bool) error {
	var number int64

	q := `SELECT COUNT(*) FROM chat_members WHERE userid = $1 AND chatid = $2`
	err := p.QueryRow(q, msg.ReceiverID, msg.ChatID).Scan(&number)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	fmt.Println(number)

	if number == 0 {
		q = `INSERT INTO chats(chatid) VALUES($1)`
		_, err = p.Exec(q, msg.ChatID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
		q = `INSERT INTO chat_members(chatid,userid) VALUES($1,$2)`
		_, err = p.Exec(q, msg.ChatID, msg.ReceiverID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}

	q = `SELECT COUNT(*) FROM chat_members WHERE userid = $1 AND chatid = $2`
	err = p.QueryRow(q, msg.SenderID, msg.ChatID).Scan(&number)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	if number == 0 {
		q = `INSERT INTO chat_members(chatid,userid) VALUES($1,$2)`
		_, err = p.Exec(q, msg.ChatID, msg.SenderID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}

	if receiverOffline {
		q = `UPDATE chat_members SET missed = missed + 1 WHERE chatid = $1 AND userid = $2`
		_, err = p.Exec(q, msg.ChatID, msg.SenderID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}

	q = `INSERT INTO messages(chatid,senderid,receiverid,messageid,val,sentat) VALUES($1,$2,$3,$4,$5,$6)`
	_, err = p.Exec(q, msg.ChatID, msg.SenderID, msg.ReceiverID, msg.MessageID, msg.Value, msg.SentAt)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Info("Database layer success")
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
	log.GetLogger(ctx).Info("Database layer success")
	return nil
}

func (p *Postgres) GetHistory(ctx context.Context, req models.GetHistoryReq, uuid string) ([]models.Message, error) {
	var msgs []models.Message
	var msg models.Message

	q := `SELECT * FROM messages WHERE chatid = $1 ORDER BY sentat DESC`
	rows, err := p.Queryx(q, req.ChatID)
	for rows.Next() {
		_ = rows.StructScan(&msg)
		msgs = append(msgs, msg)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Message{}, nil
		}
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}
	q = `UPDATE chat_members SET missed = 0 WHERE chatid = $1 AND userid = $2`
	_, err = p.Exec(q, req.ChatID, uuid)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}
	return msgs, nil
}
