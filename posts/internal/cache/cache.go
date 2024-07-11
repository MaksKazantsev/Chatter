package cache

import "context"

type Cache interface {
	Save(ctx context.Context, key string, val string) error
	Get(ctx context.Context, key string) (string, error)
}
