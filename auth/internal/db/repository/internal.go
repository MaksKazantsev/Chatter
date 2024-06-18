package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

func (p *Postgres) GetHashAndID(ctx context.Context, email string) (db.HashAndID, error) {
	var res db.HashAndID
	q := `SELECT uuid, password FROM users WHERE email = $1`

	err := p.QueryRowx(q, email).StructScan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.HashAndID{}, utils.NewError("user with this email not found", utils.ErrNotFound)
		}
		return db.HashAndID{}, utils.NewError(err.Error(), utils.ErrInternal)
	}
	return res, nil
}
