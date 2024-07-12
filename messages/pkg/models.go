package pkg

import "time"

type Message struct {
	Type string `yaml:"type"`
	Data []byte `json:"data"`
}

type UpdateOnlineMessage struct {
	ID         string    `json:"id"`
	LastOnline time.Time `json:"last_online"`
}
