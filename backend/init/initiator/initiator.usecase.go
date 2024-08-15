package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"

	employeeUsecase "fudjie.waizly/backend-test/internal/domain/employee/usecase"
)

func (i *Initiator) InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories) *service.Usecases {

	employeeUseCase := employeeUsecase.New(&employeeUsecase.Opts{
		SqlDbManager:       infra.SqldbManager,
		EmployeeRepository: repos.EmployeeRepository,
	})

	return &service.Usecases{
		EmployeeUseCase: employeeUseCase,
	}
}
