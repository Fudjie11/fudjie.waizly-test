package sqldb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

type DBEngine string

const (
	Postgres DBEngine = "postgres"
	Mysql    DBEngine = "mysql"
	Mariadb  DBEngine = "mysql"
)

type DBConfig struct {
	Driver                DBEngine      `json:"driver" yaml:"driver"`
	DSN                   string        `json:"master" yaml:"master"`
	MaxOpenConnections    int           `json:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConnections    int           `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnectionMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	Retry                 int           `json:"retry" yaml:"retry"`
}

// Master defines operation that will be executed to master DB
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)

	// ExecContext use master database to exec query
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Begin transaction on master DB
	Begin() (*sql.Tx, error)

	// BeginTx begins transaction on master DB
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Rebind a query from the default bindtype (QUESTION) to the target bindtype.
	Rebind(sql string) string

	// NamedExec do named exec on master DB
	NamedExec(query string, arg interface{}) (sql.Result, error)

	// NamedExecContext do named exec on master DB
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	// BindNamed do BindNamed on master DB
	BindNamed(query string, arg interface{}) (string, []interface{}, error)

	// Get from slave database
	Get(dest interface{}, query string, args ...interface{}) error

	// Select from slave database
	Select(dest interface{}, query string, args ...interface{}) error

	// Query from slave database
	Query(query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow executes QueryRow against slave DB
	QueryRow(query string, args ...interface{}) *sql.Row

	// NamedQuery do named query on slave DB
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	// GetContext from sql database
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// SelectContext from sql database
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// QueryContext from sql database
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext from sql database
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row

	// QueryxContext queries the database and returns an *sqlx.Rows. Any placeholder parameters are replaced with supplied args.
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)

	// QueryRowxContext queries the database and returns an *sqlx.Row. Any placeholder parameters are replaced with supplied args.
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row

	// NamedQueryContext do named query on slave DB
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
}

func Connect(ctx context.Context, cfg DBConfig) (db DB, err error) {
	dbx, err := connectWithRetry(ctx, cfg.Driver, cfg.DSN, cfg.Retry)
	if err != nil {
		return nil, err
	}

	// set conn config if there is any custom configuration
	if cfg.MaxIdleConnections > 0 {
		dbx.SetMaxIdleConns(cfg.MaxIdleConnections)
	}

	if cfg.MaxOpenConnections > 0 {
		dbx.SetMaxOpenConns(cfg.MaxOpenConnections)
	}

	if cfg.ConnectionMaxLifetime > 0 {
		dbx.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)
	}

	return dbx, nil
}

func connectWithRetry(ctx context.Context, driver DBEngine, dsn string, retry int) (*sqlx.DB, error) {
	var (
		sqlxDB *sqlx.DB
		err    error
	)

	if retry <= 0 {
		retry = 1
	}

	for i := 0; i < retry; i++ {
		sqlxDB, err = sqlx.ConnectContext(ctx, string(driver), dsn)
		if err == nil {
			return sqlxDB, nil
		}

		if i+1 < retry {
			// continue with condition
			log.Warn().Msgf("sql: retrying to connect to %s Retry: %d", dsn, i+1)
			// sleep for 3 secs in every retries
			time.Sleep(time.Second * 3)
		}
	}

	log.Error().Msgf("sql db: retry time exhausted, cannot connect to database: ", err.Error())
	err = fmt.Errorf("failed connect to database: %s", err.Error())
	return nil, err
}

func In(query string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(query, args...)
}
