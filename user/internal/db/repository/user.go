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

	if req.Username != "" {
		q = `UPDATE user_profiles SET username = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Username, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.Bio != "" {
		q = `UPDATE user_profiles SET bio = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Bio, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.Firstname != "" {
		q = `UPDATE user_profiles SET firstname = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Firstname, req.UUID)
		if err != nil {
			return utils.NewError(err.Error(), utils.ErrInternal)
		}
	}
	if req.Secondname != "" {
		q = `UPDATE user_profiles SET secondname = $1 WHERE uuid = $2`
		_, err = p.Exec(q, req.Secondname, req.UUID)
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

	log.GetLogger(ctx).Info("Database layer success")
	return nil
}

func (p *Postgres) SuggestFs(ctx context.Context, reqID, receiverID, senderID string) error {
	var UserProfileFs struct {
		Avatar     string `db:"avatar"`
		Firstname  string `db:"firstname"`
		Secondname string `db:"secondname"`
	}

	q := `SELECT avatar,firstname,secondname FROM user_profiles WHERE uuid = $1`
	err := p.QueryRowx(q, senderID).StructScan(&UserProfileFs)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	q = `INSERT INTO friend_reqs(requestid,receiverid,avatar,firstname,secondname) VALUES($1,$2,$3,$4,$5)`

	_, err = p.Exec(q, reqID, receiverID, UserProfileFs.Avatar, UserProfileFs.Firstname, UserProfileFs.Secondname)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return utils.NewError("friend req already sent", utils.ErrBadRequest)
		}
		return utils.NewError(err.Error(), utils.ErrInternal)
	}

	log.GetLogger(ctx).Info("Database layer success")
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

	log.GetLogger(ctx).Info("Database layer success")
	return nil
}

func (p *Postgres) GetFs(ctx context.Context, receiverID string) ([]models.FsReq, error) {
	var req models.FsReq
	var reqs []models.FsReq
	q := `SELECT requestid,avatar,firstname,secondname FROM friend_reqs WHERE receiverid = $1`

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

	return reqs, nil
}
