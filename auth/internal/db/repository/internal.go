package repository

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

func (p *Postgres) GetPasswordByUUID(ctx context.Context, uuid string) (string, error) {
	q := `SELECT password FROM users WHERE uuid = $1`

	var password string

	if err := p.QueryRowx(q, uuid).Scan(&password); err != nil {
		return "", utils.NewError("user with this uuid not found", utils.ErrNotFound)
	}
	return password, nil
}
