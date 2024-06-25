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
