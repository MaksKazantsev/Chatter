package broker

import (
	"context"
)

type Producer interface {
	Produce(ctx context.Context, message any)
}

type ProducerType int
