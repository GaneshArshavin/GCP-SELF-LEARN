package redis

import (
	"context"
	"time"
)

type Config struct {
	Host        string
	Port        int
	DialTimeout time.Duration
}

type Client interface {
	Get(ctx context.Context, key string) (result []byte, err error)
	Set(ctx context.Context, key string, value []byte, expiry time.Duration) error
	IncrWithTTL(ctx context.Context, key string, ttl time.Duration) (err error)
	Close()
}
