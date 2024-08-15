package service

import (
	"fudjie.waizly/backend-test/config"

	// MIDDLEWARE
	middleware "fudjie.waizly/backend-test/internal/middleware"

	employeeHandlerRest "fudjie.waizly/backend-test/internal/app/main/v1/employee/rest/handler"
	employeeRepository "fudjie.waizly/backend-test/internal/domain/employee/repository"
	employeeUseCase "fudjie.waizly/backend-test/internal/domain/employee/usecase"

	"fudjie.waizly/backend-test/library/sqldb"
)

type Infrastructure struct {
	RDBMS        sqldb.DB
	SqldbManager sqldb.SqlDbManager
}

type Repositories struct {
	employeeRepository.EmployeeRepository
}

type Usecases struct {
	employeeUseCase.EmployeeUseCase
}

type RESTHandler struct {
	employeeHandlerRest.EmployeeRestHandler
}

type RESTService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	RESTHandler    *RESTHandler
	Usecases       *Usecases
	Middleware     *Middleware
}

type Middleware struct {
	HttpMiddleware middleware.RestMiddleware
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
