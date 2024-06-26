package db

import "context"

type Repository interface {
	UploadAvatar(ctx context.Context, uuid, photoID, avatarLink string) error
	GetUserAvatar(ctx context.Context, uuid string) (string, error)
}
