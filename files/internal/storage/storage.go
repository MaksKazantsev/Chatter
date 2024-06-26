package storage

import (
	"context"
)

type Storage interface {
	Upload(ctx context.Context, id string, val []byte) (string, error)
}
