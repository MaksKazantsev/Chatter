package repository

import (
	"github.com/MaksKazantsev/SSO/auth/internal/db/migrations"
	"github.com/jmoiron/sqlx"
)

func MustConnect(addr string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic("failed to connect to db: " + err.Error())
	}
	if err = db.Ping(); err != nil {
		panic("failed to ping db: " + err.Error())
	}
	migrations.MustMigrate(db)
	return db
}
