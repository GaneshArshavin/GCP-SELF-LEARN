package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver for jmoiron/sqlx
)

type client struct {
	db          *sqlx.DB
	config      Config
}

func NewClient(config Config) Client {
	c := &client{}
	c.config = config

	log.Printf("Initializing Postgres DB: %s", c.config.Host)
	dsn := buildDSN(c.config)
	pgClient, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	c.db = pgClient
	return c
}

func (c *client) Close() {
	log.Printf("Closing Postgres DB: %s", c.config.Host)
	c.db.Close()
}

func (c *client) Exec(ctx context.Context, query string, args ...interface{}) error {

	_, err := c.db.Exec(query, args...)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := c.db.Get(dest, query, args...)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func buildDSN(c Config) string {
	var host, port string
	if strings.Contains(c.Host, ":") {
		host_port := strings.Split(c.Host, ":")
		host = host_port[0]
		port = host_port[1]
	} else {
		host = c.Host
		port = fmt.Sprintf("%d", c.Port)
	}
	return fmt.Sprintf("sslmode=disable host=%s port=%s user=%s password=%s dbname=%s", host, port, c.User, c.Password, c.DBName)
}
