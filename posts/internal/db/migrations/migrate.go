package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MustMigrate(conn *sqlx.DB) {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		panic("failed to get driver: " + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres", driver)
	if err != nil {
		panic("failed to init migrate instance: " + err.Error())
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic("failed to migrate: " + err.Error())
	}
}
