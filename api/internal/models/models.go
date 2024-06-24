package models

import "time"

type Message struct {
	ChatID     string    `json:"-"`
	SenderID   string    `json:"-"`
	ReceiverID string    `json:"receiverID"`
	MessageID  string    `json:"-"`
	Value      string    `json:"value"`
	SentAt     time.Time `json:"-"`
}
