package service

import (
	"fudjie.waizly/backend-test/config"

	tmsRoleHandlerRest "fudjie.waizly/backend-test/internal/app/main/v1/tms_role/rest/handler"
	tmsRoleRepository "fudjie.waizly/backend-test/internal/domain/tms_role/repository"
	tmsRoleUseCase "fudjie.waizly/backend-test/internal/domain/tms_role/usecase"

	// MIDDLEWARE
	middleware "fudjie.waizly/backend-test/internal/middleware"

	"fudjie.waizly/backend-test/library/sqldb"
)

type Infrastructure struct {
	RDBMS        sqldb.DB
	SqldbManager sqldb.SqlDbManager
}

type Repositories struct {
	tmsRoleRepository.TmsRoleRepository
}

type Usecases struct {
	tmsRoleUseCase.TmsRoleUseCase
}

type RESTHandler struct {
	tmsRoleHandlerRest.TmsRoleRestHandler
}

type RESTService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	RESTHandler    *RESTHandler
	Usecases       *Usecases
	Middleware     *Middleware
}

type Middleware struct {
	HttpMiddleware middleware.HttpMiddleware
}

type MessageBrokerHandler struct {
}

type MessageBrokerService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	Handlers       *MessageBrokerHandler
}

type CronHandler struct {
}

type CronService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	Handlers       *CronHandler
}
