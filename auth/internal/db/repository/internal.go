package repository

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

func (p *Postgres) GetHashAndID(ctx context.Context, email string) (string, string, error) {
	q := `SELECT password, uuid FROM users WHERE email = $1`

	var res struct {
		password string
		uuid     string
	}

	if err := p.QueryRowx(q, email).Scan(&res); err != nil {
		return "", "", utils.NewError("user with this email not found", utils.ErrNotFound)
	}
	return res.password, res.uuid, nil
}

func (p *Postgres) GetPasswordByUUID(ctx context.Context, uuid string) (string, error) {
	q := `SELECT password FROM users WHERE uuid = $1`

	var password string

	if err := p.QueryRowx(q, uuid).Scan(&password); err != nil {
		return "", utils.NewError("user with this uuid not found", utils.ErrNotFound)
	}
	return password, nil
}
