package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	"time"
)

const (
	RECOVERY     = "recovery"
	VERIFICATION = "verification"
)

func (p *Postgres) Register(ctx context.Context, req models.RegReq) error {
	q := `INSERT INTO users (uuid,email,username,password,refresh,isverified,joined) VALUES($1,$2,$3,$4,$5,$6,$7)`
	tx, err := p.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
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
	q := `INSERT INTO codes (code,email,isverified) VALUES ($1,$2,$3)`
	_, err := p.Exec(q, code, email, false)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	log.GetLogger(ctx).Debug("db layer success")
	return nil
}

func (p *Postgres) EmailVerifyCode(ctx context.Context, code, email, t string) (string, error) {
	switch t {
	case VERIFICATION:
		q := `DELETE FROM codes WHERE code = $1 AND email = $2`
		_, err := p.Exec(q, code, email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", utils.NewError("wrong code or email provided", utils.ErrNotFound)
			}
			return "", utils.NewError(err.Error(), utils.ErrInternal)
		}
	case RECOVERY:
		q := `UPDATE codes SET isverified = $1 WHERE email = $2 AND code = $3`
		_, err := p.Exec(q, true, email, code)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", utils.NewError("wrong code or email provided", utils.ErrNotFound)
			}
			return "", utils.NewError(err.Error(), utils.ErrInternal)
		}
	default:
		return "", utils.NewError("invalid verification type", utils.ErrBadRequest)
	}

	q := `UPDATE users SET isverified = $1 WHERE email = $2`
	_, err := p.Exec(q, true, email)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `SELECT uuid FROM users WHERE email = $1`
	var uuid string
	err = p.QueryRow(q, email).Scan(&uuid)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("db layer success")
	return uuid, nil
}

func (p *Postgres) PasswordRecovery(ctx context.Context, cr models.Credentials) error {
	var verified bool
	tx, err := p.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	defer tx.Rollback()
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	q := `SELECT isverified FROM codes WHERE email = $1`
	if err = tx.QueryRow(q, cr.Email).Scan(&verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user not found or code not sent", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if !verified {
		return utils.NewError("recover request not verified!", utils.ErrBadRequest)
	}
	q = `UPDATE users SET password = $1 WHERE email = $2`
	res, err := tx.Exec(q, cr.Password, cr.Email)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("user not found", utils.ErrNotFound)
	}

	q = `DELETE FROM codes WHERE email = $1 AND isverified = $2`
	res, err = tx.Exec(q, cr.Email, true)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if err = tx.Commit(); err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("db layer success")
	return nil
}
