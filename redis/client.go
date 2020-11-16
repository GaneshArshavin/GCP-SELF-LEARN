package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type client struct {
	client *redis.Client
	config Config
}

func NewClient(config Config) Client {
	r := &client{config: config}
	log.Printf("Initializing Redis: %s", r.config.Host)
	r.client = redis.NewClient(&redis.Options{
		Addr:        r.getAddr(),
		DialTimeout: r.config.DialTimeout,
	})
	return r
}

// Set sets a value into Redis with a given key.
func (r *client) Set(ctx context.Context, key string, value []byte, expiry time.Duration) (err error) {

	err = r.client.Set(key, value, expiry).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *client) Del(ctx context.Context, key string) (err error) {

	err = r.client.Del(key, key).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func (r *client) Incr(ctx context.Context, key string) (err error) {
	return r.IncrWithTTL(ctx, key, 0)
}

func (r *client) IncrWithTTL(ctx context.Context, key string, ttl time.Duration) (err error) {

	err = r.client.Incr(key).Err()
	if err != nil {
		return errors.Wrap(err, "")
	}
	// If ttl (time-to-live) is a positive number then set it as the expiry time for the key
	if ttl > 0 {
		err = r.client.Expire(key, ttl).Err()
		if err != nil {
			return errors.Wrap(err, "")
		}
	}
	if err != nil {
		return err
	}
	return nil
}

// Get gets from Redis a value for a given key. If they key does not exist,
// `redis.Nil` is returned as the error. Callers are expected to check for
// `redis.Nil` errors.
func (r *client) Get(ctx context.Context, key string) (result []byte, err error) {
	var cmd *redis.StringCmd

	cmd = r.client.Get(key)
	err = cmd.Err()
	// Return nil so hystrix doesn't register cache misses as errors
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	b, _ := cmd.Bytes()
	return b, nil
}

func (r *client) getAddr() string {
	if strings.Contains(r.config.Host, ":") {
		return r.config.Host
	}
	return fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)
}

func (r *client) Close() {
	r.client.Close()
}
