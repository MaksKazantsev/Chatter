package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

const (
	RECOVERY     = "recovery"
	VERIFICATION = "verification"
)

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
	err = p.QueryRow(q, true, email).Scan(&uuid)
	if err != nil {
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}

	return uuid, nil
}

func (p *Postgres) PasswordRecovery(ctx context.Context, cr models.Credentials) error {
	var verified bool
	tx, err := p.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	q := `SELECT isverified FROM codes WHERE email = $1`
	if err = tx.QueryRow(q, cr.Email).Scan(&verified); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user not found", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if !verified {
		return utils.NewError("recover request not verified!", utils.ErrBadRequest)
	}
	q = `UPDATE users SET password = $1 WHERE email = $2`
	res, err := tx.Exec(q, cr.Password, cr.Email)
	if err != nil {

	}
	amount, _ := res.RowsAffected()
	if amount == 0 {
		return utils.NewError("user not found", utils.ErrNotFound)
	}
	return nil
}
