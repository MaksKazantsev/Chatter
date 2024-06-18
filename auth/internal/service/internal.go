package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"math/rand"
	"strconv"
)

type Internal interface {
	EmailSendCode(ctx context.Context, email string) error
	EmailVerifyCode(ctx context.Context, code, email, t string) error
}

func (s *service) EmailVerifyCode(ctx context.Context, code, email, t string) error {
	return nil
}

func (s *service) EmailSendCode(ctx context.Context, email string) error {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// code
	code := strconv.Itoa(rand.Intn(9009) + 1000)

	// send code
	if err := s.smtp.SendCode(code, email); err != nil {
		return fmt.Errorf("smtp error: %w", err)
	}

	// calling repo method
	if err := s.repo.EmailAddCode(ctx, code, email); err != nil {
		return fmt.Errorf("repo errpr: %w", err)
	}
	return nil
}
