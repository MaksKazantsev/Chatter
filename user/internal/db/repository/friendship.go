package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
)

func (p *Postgres) SuggestFriendShip(ctx context.Context, sender, receiver string) error {
	q := `INSERT INTO friend_reqs (sender,receiver) VALUES($1,$2)`

	_, err := p.Exec(q, sender, receiver)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}

func (p *Postgres) RefuseFriendShip(ctx context.Context, sender, receiver string) error {
	q := `DELETE FROM friend_reqs WHERE sender = $1 AND receiver = $2`
	_, err := p.Exec(q, sender, receiver)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend request does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}
