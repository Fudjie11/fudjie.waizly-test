package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	tmsRoleHandlerRest "fudjie.waizly/backend-test/internal/app/main/v1/tms_role/rest/handler"
)

var ()

func (i *Initiator) InitRESTHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RESTHandler {
	tmsRoleRestHandler := tmsRoleHandlerRest.New(&tmsRoleHandlerRest.Opts{
		TmsRoleUseCase: uc.TmsRoleUseCase,
	})

	return &service.RESTHandler{
		TmsRoleRestHandler: tmsRoleRestHandler,
	}
}

func (i *Initiator) InitMessageBrokerHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.MessageBrokerHandler {
	return &service.MessageBrokerHandler{}
}

func (i *Initiator) InitCronHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.CronHandler {

	return &service.CronHandler{}
}
