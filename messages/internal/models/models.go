package models

import "time"

type Message struct {
	ChatID     string    `db:"chatid"`
	SenderID   string    `db:"senderid"`
	ReceiverID string    `db:"receiverid"`
	MessageID  string    `db:"messageid"`
	Value      string    `db:"val"`
	SentAt     time.Time `db:"sentat"`
}

type Chat struct {
	ChatID      string
	ChatPhoto   string
	ChatName    string
	Missed      int
	UserID      string
	LastMessage *Message
}

type ChatMember struct {
	UserID string
	ChatID string
	Missed int
}
