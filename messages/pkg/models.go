package pkg

import "time"

type KafkaMessage struct {
	ID         string    `json:"id"`
	LastOnline time.Time `yaml:"last_online"`
}
