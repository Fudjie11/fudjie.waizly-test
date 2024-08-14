package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	tmsRoleRepo "fudjie.waizly/backend-test/internal/domain/tms_role/repository"
)

func (i *Initiator) InitRepository(cfg *config.MainConfig, infra *service.Infrastructure) *service.Repositories {

	tmsRoleRepository := tmsRoleRepo.New(&tmsRoleRepo.Opts{
		SqlDbManager: infra.SqldbManager,
	})

	return &service.Repositories{
		TmsRoleRepository: tmsRoleRepository,
	}
}
