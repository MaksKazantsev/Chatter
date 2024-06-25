package service

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
)

type Friendship struct {
	repo db.Friendship
}

func NewFriendShip(repo db.Repository) *Friendship {
	return &Friendship{repo: repo}
}

func (f *Friendship) SuggestFriendship(ctx context.Context, req models.FriendShipReq) error {
	return nil
}

func (f *Friendship) RefuseFriendship(ctx context.Context, req models.RefuseFriendShipReq) error {
	return nil
}

func (f *Friendship) AcceptFriendship(ctx context.Context, req models.AcceptFriendShipReq) error {

	return nil
}
