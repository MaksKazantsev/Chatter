package models

import "time"

type Message struct {
	ChatID     string
	SenderID   string
	ReceiverID string
	MessageID  string
	Value      string
	SentAt     time.Time
}
