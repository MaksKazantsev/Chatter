package repository

import (
	"github.com/MaksKazantsev/Chatter/messages/internal/db"
	"github.com/jmoiron/sqlx"
)

var _ db.Repository = &Postgres{}

type Postgres struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) *Postgres {
	return &Postgres{db}
}
