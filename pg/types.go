// Package pg provides an interface to the Postgres database.
package pg

import (
	"context"
	"database/sql"
)

// Client represents the supported operations on the Postgres database. The
// methods on the client should not have service-specific business logic; they
// should only provide an interface to interact with the database.
type Client interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) error
	Close()
}

// Manager abstracts the master and slave implementation away from the service.
// The manager will perform reads on the slave and writes on the master. This
// interface inherits from the Client interface, with each manager method being
// a thin wrapper around the client's method.
// Right now going with just one
type Manager interface {
	Client
}

// Config represents a Postgres connection config
type Config struct {
	Host               string
	Port               int
	User               string
	Password           string
	DBName             string
	MaxIdleConnections int
	MaxOpenConnections int
}
