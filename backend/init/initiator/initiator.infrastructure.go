package initiator

import (
	"context"
	"time"

	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	sqlStore "fudjie.waizly/backend-test/library/sqldb"

	"fudjie.waizly/backend-test/library/sqldb"

	_ "github.com/go-sql-driver/mysql" // import mysql driver
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	errInitRdbms          = "failed to initiate RDBMS"
	errFInitMessageBroker = "failed to initiate Message Broker : %s"
	errFInitTracer        = "failed to initiate Tracer : %s"
)

var (
	rdbms        sqlStore.DB
	sqldbManager sqlStore.SqlDbManager
)

func (i *Initiator) InitInfrastructure(cfg *config.MainConfig) *service.Infrastructure {
	i.initMySQL(cfg)
	i.initRdbmsStorage()

	return &service.Infrastructure{
		RDBMS:        rdbms,
		SqldbManager: sqldbManager,
	}
}

func (i *Initiator) initMySQL(cfg *config.MainConfig) {
	var (
		newSqldb sqldb.DB
		err      error
		dbConfig = cfg.Rdbms.DBConfig
	)

	newSqldb, err = i.NewSqlDatabase(context.Background(), sqldb.DBConfig{
		Driver:                sqldb.DBEngine(dbConfig.Driver),
		DSN:                   dbConfig.DSN,
		MaxOpenConnections:    dbConfig.MaxOpenConns,
		MaxIdleConnections:    dbConfig.MaxIdleConns,
		ConnectionMaxLifetime: time.Duration(dbConfig.ConnMaxLifetime) * time.Millisecond,
		Retry:                 dbConfig.Retry,
	})

	if err != nil {
		log.Fatal().Msg(errInitRdbms)
	}

	rdbms = newSqldb
}

func (i *Initiator) initRdbmsStorage() {
	sqldbManager = sqlStore.New(&sqlStore.Opts{
		DB: rdbms,
	})
}
