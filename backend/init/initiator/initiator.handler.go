package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	employeeHandlerRest "fudjie.waizly/backend-test/internal/app/main/v1/employee/rest/handler"
)

var ()

func (i *Initiator) InitRESTHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RESTHandler {
	employeeRestHandler := employeeHandlerRest.New(&employeeHandlerRest.Opts{
		EmployeeUseCase: uc.EmployeeUseCase,
	})

	return &service.RESTHandler{
		EmployeeRestHandler: employeeRestHandler,
	}
}

func (i *Initiator) InitMessageBrokerHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.MessageBrokerHandler {
	return &service.MessageBrokerHandler{}
}

func (i *Initiator) InitCronHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.CronHandler {

	return &service.CronHandler{}
}
