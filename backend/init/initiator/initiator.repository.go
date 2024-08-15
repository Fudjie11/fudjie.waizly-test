package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	employeeRepo "fudjie.waizly/backend-test/internal/domain/employee/repository"
)

func (i *Initiator) InitRepository(cfg *config.MainConfig, infra *service.Infrastructure) *service.Repositories {

	employeeRepository := employeeRepo.New(&employeeRepo.Opts{
		SqlDbManager: infra.SqldbManager,
	})

	return &service.Repositories{
		EmployeeRepository: employeeRepository,
	}
}
