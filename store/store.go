package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/carousell/chope-assignment/model"
	"github.com/carousell/chope-assignment/pg"
	"github.com/carousell/chope-assignment/redis"
	"github.com/pkg/errors"
)

const (
	AccountsUserTable     = "accountsuser"
	AccountsActivityTable = "accounts_activity"
	httpAddress           = "http://127.0.0.1:9282/"
)

type storage struct {
	db    pg.Manager
	redis redis.Client
	psql  sq.StatementBuilderType
}

func NewClient(pgMasterConfig *pg.Config, pgSlaveConfig *pg.Config, redisConfig *redis.Config) (StorageService, error) {
	s := new(storage)
	s.db = pg.NewManager(*pgMasterConfig, *pgSlaveConfig)
	s.redis = redis.NewClient(*redisConfig)
	s.psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return s, nil
}

func (s *storage) GetLoginAttempts(ctx context.Context, username string) (count int, err error) {

	err = s.RedisGet(ctx, redis.GetloginAttemptsKey(username), &count)
	if err != nil {
		return count, errors.Wrap(err, "Storage Error : Erorr fetching count from redis")
	}
	fmt.Println("count", count)
	return count, err
}

func (s *storage) GetAccountsUser(ctx context.Context, username string) ([]*model.AccountsUser, error) {
	users := []*model.AccountsUser{}
	query := fmt.Sprintf("Select id,username,email,created_at,updated_at,passowrd from accountsuser where username = '%s';", username)
	rows, err := s.db.Query(ctx, query, nil)
	for rows.Next() {
		b := &model.AccountsUser{}
		if err = rows.Scan(&b.ID, &b.Username, &b.Email, &b.CreatedAt, &b.UpdatedAt, &b.Passowrd); err != nil {
			fmt.Println("err", err)
		}
		users = append(users, b)
	}
	return users, nil
}

func (s *storage) RedisGet(ctx context.Context, key string, dest interface{}) error {
	b, err := s.redis.Get(ctx, key)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = json.Unmarshal(b, dest)
	if err != nil {
		err = errors.Wrap(err, "cannot unmarshal value from redis")
		return err
	}
	return nil
}

func (s *storage) CreateUser(ctx context.Context, m *model.AccountsUser) (err error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s','%s','%s')", AccountsUserTable, "\"username\",\"passowrd\",\"email\"", m.Username.String, m.Passowrd.String, m.Email.String)
	err = s.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) IncrementRedisRetryCounter(ctx context.Context, Username string) (err error) {
	ttl, _ := time.ParseDuration("2m")
	err = s.redis.IncrWithTTL(ctx, redis.GetloginAttemptsKey(Username), ttl)
	if err != nil {
		return err
	}
	return nil
}
func (s *storage) StoreLoginActivity(ctx context.Context, userID string, company string, isSuccess bool) (err error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s','%s',%t,'%s')", AccountsActivityTable, "\"user_id\",\"company_name\",\"is_success\",\"operation_type\"", userID, company, isSuccess, "Login")
	err = s.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
func (s *storage) StoreInHouseToken(ctx context.Context, token string, userID string, duration string) (string, error) {
	expiry, err := time.ParseDuration(duration)
	if err != nil {
		return "", errors.New("Redis Error : Invalid time duration")
	}
	err = s.redis.Set(ctx, redis.GetInHouseTokenKey(token), []byte(userID), expiry)
	if err != nil {
		return "", errors.New("Redis Error : Erorr in setting In house token")
	}
	return redis.GetInHouseTokenKey(token), nil
}

func (s *storage) StoreInThirdPartyToken(ctx context.Context, token string, userID string, duration string) error {
	expiry, err := time.ParseDuration(duration)
	if err != nil {
		return errors.New("Redis Error : Invalid time duration")
	}
	err = s.redis.Set(ctx, redis.GetThirdPartyTokenKey(token), []byte(userID), expiry)
	if err != nil {
		return errors.New("Redis Error : Erorr in setting Thrid party token")
	}
	return nil
}

func (s *storage) Close() {
	s.db.Close()
}
