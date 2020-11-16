package store

import (
	"context"

	"github.com/carousell/chope-assignment/model"
)

type StorageService interface {
	CreateUser(ctx context.Context, m *model.AccountsUser) error
	GetAccountsUser(ctx context.Context, username string) (user []*model.AccountsUser, err error)
	GetLoginAttempts(ctx context.Context, username string) (int, error)
	StoreInHouseToken(ctx context.Context, token string, userID string, duration string) (string, error)
	StoreInThirdPartyToken(ctx context.Context, token string, userID string, duration string) error
	StoreLoginActivity(ctx context.Context, userID string, company string, is_success bool) error
	IncrementRedisRetryCounter(ctx context.Context, userID string) error
	Close()
}
