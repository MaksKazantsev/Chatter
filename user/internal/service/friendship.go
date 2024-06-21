package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/Chatter/user/internal/db"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
)

type Friendship struct {
	repo db.Friendship
}

func (f *Friendship) SuggestFriendship(ctx context.Context, token, receiver string) error {
	log.GetLogger(ctx).Debug("uc layer success ✔")
	cl, err := utils.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if err = f.repo.SuggestFriendShip(ctx, cl["id"].(string), receiver); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (f *Friendship) RefuseFriendship(ctx context.Context, token, receiver string) error {
	log.GetLogger(ctx).Debug("uc layer success ✔")
	cl, err := utils.ParseToken(token)
	if err != nil {
		return fmt.Errorf("failed to parse token: %w", err)
	}
	if err = f.repo.RefuseFriendShip(ctx, cl["id"].(string), receiver); err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}

func (f *Friendship) AcceptFriendship(ctx context.Context, receiver string, sender string) error {
	//TODO implement me
	panic("implement me")
}

func NewFriendShip(repo db.Repository) *Friendship {
	return &Friendship{repo: repo}
}
