package cache

import (
	"context"
	"github.com/MaksKazantsev/Chatter/posts/internal/utils"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
	"time"
)

type redisCL struct {
	cl *redis.Client
}

func NewRedis() Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	return &redisCL{cl: rdb}
}

func (c *redisCL) Get(ctx context.Context, key string) (string, error) {
	res, err := c.cl.Get(ctx, key).Result()
	if err != nil {
		if strings.Contains(err.Error(), "redis: nil") {
			return "", nil
		}
		return "", utils.NewError(err.Error(), utils.ErrInternal)
	}
	return res, nil
}
func (c *redisCL) Save(ctx context.Context, key string, val string) error {
	if err := c.cl.Set(ctx, key, val, time.Minute*15).Err(); err != nil {
		return utils.NewError("failed to cache", utils.ErrInternal)
	}
	return nil
}
