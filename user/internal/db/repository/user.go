package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/MaksKazantsev/Chatter/user/internal/log"
	"github.com/MaksKazantsev/Chatter/user/internal/models"
	"github.com/MaksKazantsev/Chatter/user/internal/utils"
	"strings"
	"time"
)

func (p *Postgres) EditProfile(ctx context.Context, req models.UserProfile) error {
	var q string
	q = `SELECT username FROM users WHERE uuid = $1`
	_, err := p.Exec(q, req.UUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user with this uuid does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	if req.Bio != "" {
		q = `UPDATE user_profiles SET bio = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Bio, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.Avatar != "" {
		q = `UPDATE user_profiles SET avatar = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Avatar, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if !req.Birthday.IsZero() {
		q = `UPDATE user_profiles SET birthday = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Birthday, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}

	q = `UPDATE user_profiles SET lastonline = $1 WHERE uuid = $2`
	_, err = p.Exec(q, time.Now(), req.UUID)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) SuggestFs(ctx context.Context, reqID, receiverID, senderID string) error {
	var UserProfileFs struct {
		Avatar   string `db:"avatar"`
		Username string `db:"username"`
	}

	q := `SELECT avatar,username FROM user_profiles WHERE uuid = $1`
	err := p.QueryRowx(q, senderID).StructScan(&UserProfileFs)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `INSERT INTO friend_reqs(requestid,receiverid,avatar,username) VALUES($1,$2,$3,$4)`

	_, err = p.Exec(q, reqID, receiverID, UserProfileFs.Avatar, UserProfileFs.Username)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.NewError("friend req already sent", utils.ErrBadRequest)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) RefuseFs(ctx context.Context, reqID, receiverID string) error {
	q := `DELETE FROM friend_reqs WHERE requestid = $1 AND receiverid = $2`
	_, err := p.Exec(q, reqID, receiverID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend req does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) GetFs(ctx context.Context, receiverID string) ([]models.FsReq, error) {
	var req models.FsReq
	var reqs []models.FsReq
	q := `SELECT requestid,avatar,username FROM friend_reqs WHERE receiverid = $1`

	rows, err := p.Queryx(q, receiverID)
	for rows.Next() {
		_ = rows.StructScan(&req)
		reqs = append(reqs, req)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.FsReq{}, nil
		}
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return reqs, nil
}

func (p *Postgres) AcceptFs(ctx context.Context, senderID, receiverID, reqID string) error {
	q := `SELECT avatar,username FROM friend_reqs WHERE receiverid = $1`
	var Fs struct {
		Avatar   string `db:"avatar"`
		Username string `db:"username"`
	}

	if err := p.QueryRowx(q, receiverID).StructScan(&Fs); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend req does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `INSERT INTO friends(uuid,friendid,avatar,username) VALUES($1,$2,$3,$4)`
	_, err := p.Exec(q, receiverID, senderID, Fs.Avatar, Fs.Username)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.NewError("friend already exists", utils.ErrBadRequest)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `SELECT avatar,username FROM user_profiles WHERE uuid = $1`
	if err = p.QueryRowx(q, receiverID).StructScan(&Fs); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend req does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `INSERT INTO friends(uuid,friendid,avatar,username) VALUES($1,$2,$3,$4)`
	_, err = p.Exec(q, senderID, receiverID, Fs.Avatar, Fs.Username)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.NewError("friend already exists", utils.ErrBadRequest)
		}
	}

	q = `DELETE FROM friend_reqs WHERE requestid = $1 AND receiverid = $2`
	_, err = p.Exec(q, reqID, receiverID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend req does not exist", utils.ErrBadRequest)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) DeleteFriend(ctx context.Context, userID, friendID string) error {
	q := `DELETE FROM friends WHERE uuid = $1 AND friendid = $2 OR uuid = $3 AND friendid = $4`
	_, err := p.Exec(q, userID, friendID, friendID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("friend does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}

func (p *Postgres) GetFriends(ctx context.Context, userID string) ([]models.Friend, error) {
	var friend models.Friend
	var friends []models.Friend
	q := `SELECT friendid,avatar,username FROM friends WHERE uuid = $1`

	rows, err := p.Queryx(q, userID)
	for rows.Next() {
		_ = rows.StructScan(&friend)
		friends = append(friends, friend)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.ErrInternal)
	}
	return friends, nil
}

func (p *Postgres) GetProfile(ctx context.Context, userID string) (models.GetUserProfile, error) {
	var profile models.GetUserProfile
	q := `SELECT avatar,username,bio,birthday,lastonline FROM user_profiles WHERE uuid = $1`

	err := p.QueryRowx(q, userID).StructScan(&profile)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetUserProfile{}, utils.NewError("user does not exist", utils.ErrNotFound)
		}
		return models.GetUserProfile{}, utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Debug("Database layer success")
	return profile, nil
}

func (p *Postgres) EditAvatar(ctx context.Context, userID, avatar string) error {
	q := `UPDATE user_profiles SET avatar = $1 WHERE uuid = $2`

	_, err := p.Exec(q, avatar, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	log.GetLogger(ctx).Debug("Database layer success")
	return nil

}

func (p *Postgres) DeleteAvatar(ctx context.Context, userID string) error {
	q := `UPDATE user_profiles SET avatar = '' WHERE uuid = $1`
	_, err := p.Exec(q, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.NewError("user does not exist", utils.ErrNotFound)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	log.GetLogger(ctx).Debug("Database layer success")
	return nil
}
