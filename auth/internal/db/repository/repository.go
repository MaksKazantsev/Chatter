package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	*sqlx.DB
}

var _ db.Repository = &Postgres{}

func NewRepository(db *sqlx.DB) *Postgres {
	return &Postgres{
		db,
	}
}

func (p *Postgres) Register(ctx context.Context, req models.RegReq) error {
	q := `INSERT INTO users (uuid,email,username,password,refresh) VALUES($1,$2,$3,$4,$5)`

	_, err := p.Exec(q, req.UUID, req.Email, req.Username, req.Password, req.Refresh)
	if err != nil {
		return utils.NewError("user with this email already exists", utils.ErrBadRequest)
	}

	log.GetLogger(ctx).Debug("db layer success")
	return nil
}

func (p *Postgres) Login(ctx context.Context, req models.LogReq) error {
	q := `UPDATE users SET refresh = $1 WHERE email = $2`

	_, err := p.Exec(q, req.Refresh, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user with this email not found", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("db layer success")
	return nil
}

func (p *Postgres) Reset(ctx context.Context, password, uuid string) error {
	q := `UPDATE users SET password = $1 WHERE uuid = $2`

	_, err := p.Exec(q, password, uuid)
	if err != nil {
		return utils.NewError("user with this id not found", utils.ErrNotFound)
	}
	log.GetLogger(ctx).Debug("repo layer success ✔")

	return nil
}
