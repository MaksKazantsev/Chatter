package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
)

type Verification struct {
	repo db.Verification
}

func NewVerification(repo db.Verification) *Verification {
	return &Verification{
		repo: repo,
	}
}

func (v *Verification) EmailVerifyCode(ctx context.Context, code, email, t string) (string, string, error) {
	id, err := v.repo.EmailVerifyCode(ctx, code, email, t)
	if err != nil {
		return "", "", fmt.Errorf("repo error: %w", err)
	}
	aToken, err := utils.NewToken(id, utils.ACCESS)
	if err != nil {
		return "", "", fmt.Errorf("failed to create token: %w", err)
	}
	rToken, err := utils.NewToken(id, utils.REFRESH)
	if err != nil {
		return "", "", fmt.Errorf("failed to create token: %w", err)
	}
	if err = v.repo.UpdateRToken(ctx, id, rToken); err != nil {
		return "", "", fmt.Errorf("repo error: %w", err)
	}
	return aToken, rToken, nil
}

func (v *Verification) PasswordRecovery(ctx context.Context, cr models.Credentials) error {
	// hashing password
	hashed, err := utils.HashPassword(cr.Password)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	cr.Password = hashed

	err = v.repo.PasswordRecovery(ctx, cr)
	if err != nil {
		return fmt.Errorf("repo error: %w", err)
	}
	return nil
}
