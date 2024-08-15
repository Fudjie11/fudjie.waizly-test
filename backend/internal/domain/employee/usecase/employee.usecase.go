package usecase

import (
	"context"

	"github.com/google/uuid"

	"fudjie.waizly/backend-test/internal/db/entity"
	model "fudjie.waizly/backend-test/internal/domain/employee/model"
	employeeRepo "fudjie.waizly/backend-test/internal/domain/employee/repository"

	iModel "fudjie.waizly/backend-test/internal/model"
	sqlStore "fudjie.waizly/backend-test/library/sqldb"
)

type EmployeeUseCase interface {
	GetEmployees(ctx context.Context, pg iModel.PaginationAndSearch) (*iModel.ListPaging[model.EmployeeDto], error)
	GetEmployeeByEmployeeId(ctx context.Context, branchId uuid.UUID) (*entity.Employee, error)

	CreateEmployee(ctx context.Context, dto *model.AddEmployeeDto) (uuid.UUID, error)
	UpdateEmployee(ctx context.Context, dto *model.UpdateEmployeeDto) error
	DeleteEmployee(ctx context.Context, employeeId uuid.UUID) error
}

type Module struct {
	sqlDbManager       sqlStore.SqlDbManager
	employeeRepository employeeRepo.EmployeeRepository
}

type Opts struct {
	SqlDbManager       sqlStore.SqlDbManager
	EmployeeRepository employeeRepo.EmployeeRepository
}

func New(o *Opts) EmployeeUseCase {
	return &Module{
		sqlDbManager:       o.SqlDbManager,
		employeeRepository: o.EmployeeRepository,
	}
}
