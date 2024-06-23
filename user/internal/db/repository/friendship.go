package repository

import (
	"context"
)

func (p *Postgres) SuggestFriendShip(ctx context.Context, sender, receiver string) error {
	return nil
}

func (p *Postgres) RefuseFriendShip(ctx context.Context, reqAccepter, reqSender string) error {
	return nil
}
