package store

import (
	"context"

	"github.com/carousell/chope-assignment/model"
)

type StorageService interface {
	CreateUser(ctx context.Context, m *model.AccountsUser) error
	Close()
}
