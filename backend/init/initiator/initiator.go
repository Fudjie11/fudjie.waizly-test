package initiator

import (
	"flag"

	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	configpkg "fudjie.waizly/backend-test/library/config"
	"fudjie.waizly/backend-test/library/sqldb"

	"context"

	"github.com/google/gops/agent"
)

type Init func()
type agentListen func(opts agent.Options) error
type ReadConfig func(cfg interface{}, path string, module string) error
type NewSqlDatabase func(ctx context.Context, cfg sqldb.DBConfig) (db sqldb.DB, err error)

type InitiatorManager interface {
	InitConfig(configPath string, serviceName string) *config.MainConfig
	InitInfrastructure(cfg *config.MainConfig) *service.Infrastructure
	InitRepository(cfg *config.MainConfig, infra *service.Infrastructure) *service.Repositories
	InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories) *service.Usecases
	InitRESTHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RESTHandler
	InitRESTService(cfg *config.MainConfig, infra *service.Infrastructure, middleware *service.Middleware, hdl *service.RESTHandler, uc *service.Usecases) *service.RESTService
	InitMiddleware(cfg *config.MainConfig) *service.Middleware
}

type Initiator struct {
	FlagParse      Init
	AgentListen    agentListen
	ReadConfig     ReadConfig
	NewSqlDatabase NewSqlDatabase
}

func New() InitiatorManager {
	return &Initiator{
		FlagParse:      flag.Parse,
		AgentListen:    agent.Listen,
		ReadConfig:     configpkg.ReadConfig,
		NewSqlDatabase: sqldb.Connect,
	}
}

func (i *Initiator) InitRESTService(cfg *config.MainConfig, infra *service.Infrastructure, middleware *service.Middleware, hdl *service.RESTHandler, uc *service.Usecases) *service.RESTService {
	svc := service.RESTService{
		Config:         cfg,
		Infrastructure: infra,
		RESTHandler:    hdl,
		Middleware:     middleware,
	}

	return &svc
}

func (i *Initiator) InitMessageBrokerService(cfg *config.MainConfig, infra *service.Infrastructure, hdl *service.MessageBrokerHandler, uc *service.Usecases) *service.MessageBrokerService {
	svc := service.MessageBrokerService{
		Config:         cfg,
		Infrastructure: infra,
		Handlers:       hdl,
	}

	return &svc
}

func (i *Initiator) InitCronService(cfg *config.MainConfig, infra *service.Infrastructure, hdl *service.CronHandler) *service.CronService {
	svc := service.CronService{
		Config:         cfg,
		Infrastructure: infra,
		Handlers:       hdl,
	}

	return &svc
}
