package models

import "time"

type Message struct {
	ChatID     string    `json:"chatID"`
	SenderID   string    `json:"senderID"`
	ReceiverID string    `json:"receiverID"`
	MessageID  string    `json:"messageID"`
	Value      string    `json:"value"`
	SentAt     time.Time `json:"sentAt"`
	Token      string    `json:"token"`
}
