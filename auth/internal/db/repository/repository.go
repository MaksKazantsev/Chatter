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
	"time"
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
	q := `INSERT INTO users (uuid,email,username,password,refresh,isverified,join) VALUES($1,$2,$3,$4,$5,$6,$7)`
	tx, err := p.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	_, err = tx.Exec(q, req.UUID, req.Email, req.Username, req.Password, req.Refresh, false, time.Now())
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `INSERT INTO user_profiles (uuid,username,email,birthday,bio,lastonline,firstname,secondname) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err = tx.Exec(q, req.UUID, req.Username, req.Email, " ", " ", time.Now(), " ", " ")
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	if err = tx.Commit(); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
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

	q = `UPDATE user_profiles SET lastonline = $1 WHERE email = $2`
	_, err = p.Exec(q, time.Now(), req.Email)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("db layer success")
	return nil
}

func (p *Postgres) EmailAddCode(ctx context.Context, code, email string) error {
	q := `INSERT INTO codes (code,email) VALUES ($1,$2)`
	_, err := p.Exec(q, code, email)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}

func (p *Postgres) UpdateRToken(ctx context.Context, id, rToken string) error {
	q := `UPDATE users SET refresh = $1 WHERE uuid = $2`
	_, err := p.Exec(q, rToken, id)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}
