package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/SSO/auth/internal/db"
	"github.com/MaksKazantsev/SSO/auth/internal/log"
	"github.com/MaksKazantsev/SSO/auth/internal/models"
	"github.com/MaksKazantsev/SSO/auth/internal/utils"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

type Auth struct {
	repo db.Auth
	smtp utils.Smtp
}

func NewAuth(repo db.Auth) *Auth {
	return &Auth{
		repo: repo,
		smtp: utils.NewSmtp(),
	}
}

func (a *Auth) Register(ctx context.Context, req models.RegReq) (models.RegRes, error) {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// hashing password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.RegRes{}, err
	}
	req.Password = hashed

	// generating id
	req.UUID = uuid.New().String()

	// generating tokens
	rToken, err := utils.NewToken(req.UUID, utils.REFRESH)
	if err != nil {
		return models.RegRes{}, err
	}
	aToken, err := utils.NewToken(req.UUID, utils.ACCESS)
	if err != nil {
		return models.RegRes{}, err
	}
	req.Refresh = rToken

	if err = a.repo.Register(ctx, req); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}

	code := strconv.Itoa(rand.Intn(9001) + 1000)

	go func() {
		if err = a.smtp.SendCode(code, req.Email); err != nil {
			fmt.Printf("smtp error: %v\n", err)
		}
	}()
	if err = a.repo.EmailAddCode(ctx, code, req.Email); err != nil {
		return models.RegRes{}, fmt.Errorf("repo error: %w", err)
	}
	return models.RegRes{
		UUID:         req.UUID,
		RefreshToken: rToken,
		AccessToken:  aToken,
	}, nil
}

func (a *Auth) Login(ctx context.Context, req models.LogReq) (string, string, error) {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// get password
	res, err := a.repo.GetHashAndID(ctx, req.Email)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	// compare passwords
	if err = utils.ComparePass(res.Password, req.Password); err != nil {
		return "", " ", fmt.Errorf("error: %w", err)
	}

	// generating token
	rToken, err := utils.NewToken(res.UUID, utils.REFRESH)
	if err != nil {
		return " ", " ", err
	}
	aToken, err := utils.NewToken(res.UUID, utils.ACCESS)
	if err != nil {
		return " ", " ", err
	}
	req.Refresh = rToken

	// login
	err = a.repo.Login(ctx, req)
	if err != nil {
		return "", " ", fmt.Errorf("repo error: %w", err)
	}

	return rToken, aToken, nil
}

func (a *Auth) EmailSendCode(ctx context.Context, email string) error {
	// logging
	log.GetLogger(ctx).Debug("uc layer success ✔")

	// code
	code := strconv.Itoa(rand.Intn(9009) + 1000)

	// send code
	if err := a.smtp.SendCode(code, email); err != nil {
		return fmt.Errorf("smtp error: %w", err)
	}

	// calling repo method
	if err := a.repo.EmailAddCode(ctx, code, email); err != nil {
		return fmt.Errorf("repo errpr: %w", err)
	}
	return nil
}
