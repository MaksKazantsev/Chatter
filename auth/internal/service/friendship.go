package service

import (
	"context"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
)

type FriendShip interface {
	SuggestFriendship(ctx context.Context, sender string, receiver string) error
	RefuseFriendship(ctx context.Context, receiver string, sender string) error
	AcceptFriendship(ctx context.Context, receiver string, sender string) error
}

type friendship struct {
	repo db.Repository
}

func (f friendship) SuggestFriendship(ctx context.Context, sender string, receiver string) error {
	//TODO implement me
	panic("implement me")
}

func (f friendship) RefuseFriendship(ctx context.Context, receiver string, sender string) error {
	//TODO implement me
	panic("implement me")
}

func (f friendship) AcceptFriendship(ctx context.Context, receiver string, sender string) error {
	//TODO implement me
	panic("implement me")
}

func NewFriendShip(repo db.Repository) FriendShip {
	return &friendship{repo: repo}
}
