package store

import (
	"context"
	"fmt"

	"github.com/carousell/chope-assignment/model"
	"github.com/carousell/chope-assignment/pg"
)

const (
	AccountsUserTable = "accounts_user"
	httpAddress       = "http://127.0.0.1:9282/"
)

type storage struct {
	db pg.Manager
}

func NewClient(pgMasterConfig *pg.Config, pgSlaveConfig *pg.Config) (StorageService, error) {
	s := new(storage)
	s.db = pg.NewManager(*pgMasterConfig, *pgSlaveConfig)
	return s, nil
}

func (s *storage) CreateUser(ctx context.Context, m *model.AccountsUser) (err error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s','%s','%s')", AccountsUserTable, "\"username\",\"passowrd\",\"email\"", m.Username.String, m.Passowrd.String, m.Email.String)
	err = s.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Close() {
	s.db.Close()
}
