package pg

import (
	"context"
	"database/sql"
)

type manager struct {
	dbMaster Client
	dbSlave  Client
}

func NewManager(dbMasterConfig Config, dbSlaveConfig Config) Manager {
	m := &manager{}
	m.dbMaster = NewClient(dbMasterConfig)
	m.dbSlave = NewClient(dbSlaveConfig)
	return m
}

func (m *manager) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.dbSlave.Get(ctx, dest, query, args...)
}

func (m *manager) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return m.dbMaster.Query(ctx, query, args...)
}

func (m *manager) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return m.dbMaster.Select(ctx, dest, query, args...)
}

func (m *manager) Exec(ctx context.Context, query string, args ...interface{}) error {
	return m.dbMaster.Exec(ctx, query, args...)
}

func (m *manager) GetMaster() Client {
	return m.dbMaster
}

func (m *manager) GetSlave() Client {
	return m.dbSlave
}

func (m *manager) Close() {
	m.dbMaster.Close()
	m.dbSlave.Close()
}
