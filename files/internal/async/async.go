package async

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, message any)
}
