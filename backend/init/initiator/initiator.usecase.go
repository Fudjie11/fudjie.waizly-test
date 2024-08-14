package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	tmsRoleUsecase "fudjie.waizly/backend-test/internal/domain/tms_role/usecase"
)

func (i *Initiator) InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories) *service.Usecases {

	tmsRoleUseCase := tmsRoleUsecase.New(&tmsRoleUsecase.Opts{
		SqlDbManager:      infra.SqldbManager,
		TmsRoleRepository: repos.TmsRoleRepository,
	})

	return &service.Usecases{
		TmsRoleUseCase: tmsRoleUseCase,
	}
}
