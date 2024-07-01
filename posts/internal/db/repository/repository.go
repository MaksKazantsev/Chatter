package repository

import (
	"github.com/MaksKazantsev/Chatter/posts/internal/db"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	*sqlx.DB
}

var _ db.Repository = &Postgres{}

func NewRepository(conn *sqlx.DB) *Postgres {
	return &Postgres{conn}
}
