package repository

import (
	"context"
	"github.com/MaksKazantsev/Chatter/files/internal/db"
	"github.com/MaksKazantsev/Chatter/files/internal/utils"
	"github.com/jmoiron/sqlx"
)

var _ db.Repository = &Postgres{}

type Postgres struct {
	*sqlx.DB
}

func (p *Postgres) UploadAvatar(ctx context.Context, uuid, photoID, avatarLink string) error {
	q := `INSERT INTO files(filelink,fileid,userid) VALUES($1,$2,$3)`

	_, err := p.Exec(q, avatarLink, photoID, uuid)
	if err != nil {
		return utils.NewError(err.Error(), utils.ErrInternal)
	}
	return nil
}

func (p *Postgres) GetUserAvatar(ctx context.Context, uuid string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(conn *sqlx.DB) *Postgres {
	return &Postgres{
		conn,
	}
}
