package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	entity "fudjie.waizly/backend-test/internal/db/entity"
	model "fudjie.waizly/backend-test/internal/domain/employee/model"
	iModel "fudjie.waizly/backend-test/internal/model"
)

func (m *Module) GetEmployees(ctx context.Context, pg iModel.PaginationAndSearch) (*iModel.ListPaging[model.EmployeeDto], error) {
	var (
		err       error
		results   []model.EmployeeDto
		totalRows int
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		results, err = m.employeeRepository.GetEmployees(gctx, pg)
		return err
	})

	g.Go(func() error {
		totalRows, err = m.employeeRepository.GetEmployeeTotalRows(gctx, pg)
		return err
	})

	if err = g.Wait(); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	resultData := iModel.NewListPaging(results, int32(totalRows), pg.Limit, pg.Offset)

	return &resultData, nil
}

func (m *Module) GetEmployeeByEmployeeId(ctx context.Context, employeeId uuid.UUID) (*entity.Employee, error) {
	var (
		eEmployee *entity.Employee
		err       error
	)

	eEmployee, err = m.employeeRepository.GetEntityEmployeeByEmployeeId(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	return eEmployee, nil
}
