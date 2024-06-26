package models

import "time"

type Photo struct {
	PhotoID     string
	CreatorID   string
	CreatorName string
	PhotoLink   string
	CreatedAt   time.Time
}
