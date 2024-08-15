package repository

import (
	"context"

	"github.com/google/uuid"

	"fudjie.waizly/backend-test/internal/db/entity"
	"fudjie.waizly/backend-test/internal/domain/employee/model"
	iModel "fudjie.waizly/backend-test/internal/model"
	sqlStore "fudjie.waizly/backend-test/library/sqldb"
)

type EmployeeRepository interface {
	GetEmployees(ctx context.Context, pg iModel.PaginationAndSearch) ([]model.EmployeeDto, error)
	GetEmployeeTotalRows(ctx context.Context, pg iModel.PaginationAndSearch) (int, error)
	GetEntityEmployeeByEmployeeId(ctx context.Context, employeeId uuid.UUID) (*entity.Employee, error)

	CreateEmployee(ctx context.Context, data *entity.Employee) error
	UpdateEmployee(ctx context.Context, data *entity.Employee) (int64, error)
	DeleteEmployee(ctx context.Context, employeeId uuid.UUID) (rowsAffected int64, err error)
}

type Module struct {
	sqlDbManager sqlStore.SqlDbManager
}

type Opts struct {
	SqlDbManager sqlStore.SqlDbManager
}

func New(o *Opts) EmployeeRepository {
	return &Module{
		sqlDbManager: o.SqlDbManager,
	}
}
