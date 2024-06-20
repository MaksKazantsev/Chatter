package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
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

func (p *Postgres) UpdateRToken(ctx context.Context, id, rToken string) error {
	q := `UPDATE users SET refresh = $1 WHERE uuid = $2`
	_, err := p.Exec(q, rToken, id)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}
