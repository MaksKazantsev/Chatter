package service

import (
	"context"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
)

type Friendship struct {
	repo db.Repository
}

func (f *Friendship) SuggestFriendship(ctx context.Context, sender string, receiver string) error {
	//TODO implement me
	panic("implement me")
}

func (f *Friendship) RefuseFriendship(ctx context.Context, receiver string, sender string) error {
	//TODO implement me
	panic("implement me")
}

func (f *Friendship) AcceptFriendship(ctx context.Context, receiver string, sender string) error {
	//TODO implement me
	panic("implement me")
}

func NewFriendShip(repo db.Repository) *Friendship {
	return &Friendship{repo: repo}
}
