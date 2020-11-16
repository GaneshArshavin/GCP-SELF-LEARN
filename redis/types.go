package redis

import "time"

type Config struct {
	Host        string
	Port        int
	DialTimeout time.Duration
}

type Client interface {
	Close()
}
