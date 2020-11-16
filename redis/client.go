package redis

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis"
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

func (r *client) getAddr() string {
	if strings.Contains(r.config.Host, ":") {
		return r.config.Host
	}
	return fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)
}

func (r *client) Close() {
	r.client.Close()
}
