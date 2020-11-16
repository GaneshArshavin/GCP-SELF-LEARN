package store

import (
	"context"

	"github.com/carousell/chope-assignment/model"
)

type StorageService interface {
	CreateUser(ctx context.Context, m *model.AccountsUser) (*model.AccountsUser, error)
	GetUsersByUsernameOrEmail(ctx context.Context, username string, email string) (user []*model.AccountsUser, err error)
	GetLoginAttempts(ctx context.Context, username string) (int, error)
	StoreInHouseToken(ctx context.Context, token string, userID string, duration string) (string, error)
	GetInHouseToken(ctx context.Context, token string) (string, error)
	StoreInThirdPartyToken(ctx context.Context, token string, userID string, duration string, companyName string) error
	RemoveInHouseToken(ctx context.Context, token string) error
	StoreActivity(ctx context.Context, userID string, company string, is_success bool, activity string) error
	IncrementRedisRetryCounter(ctx context.Context, userID string) error
	Close()
}
