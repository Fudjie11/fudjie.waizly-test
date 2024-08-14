package sqldb

import (
	"context"
	"database/sql"
)

type WrapTransactionFunc = func(ctx context.Context, tx *sql.Tx) error

type module struct {
	db DB
}

type Opts struct {
	DB DB
}

func New(o *Opts) SqlDbManager {
	return &module{
		db: o.DB,
	}
}

type SqlDbManager interface {
	Store() DB
	StartTransaction(ctx context.Context) (*sql.Tx, error)
}

func (m *module) Store() DB {
	return m.db
}

func (m *module) StartTransaction(ctx context.Context) (*sql.Tx, error) {
	return m.db.BeginTx(ctx, nil)
}
