package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"fudjie.waizly/backend-test/internal/db/entity"
	"fudjie.waizly/backend-test/internal/domain/employee/model"
	"fudjie.waizly/backend-test/internal/helper"

	iModel "fudjie.waizly/backend-test/internal/model"
)

func (m *Module) GetEmployees(ctx context.Context, pg iModel.PaginationAndSearch) ([]model.EmployeeDto, error) {
	var (
		resultData = []model.EmployeeDto{}
		sb         strings.Builder
	)

	q := `select
			employee_id,
			name,
			job_title,
			salary,
			department,
			joined_date
		from public.employee`

	sb.WriteString(q)
	sb.WriteString(pg.BuildPaginationAndSearchQuery(true))

	qResult := m.sqlDbManager.Store().Rebind(sb.String())
	err := m.sqlDbManager.Store().SelectContext(ctx, &resultData, qResult, pg.Search)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, helper.NewSqlErr(err)
	}

	return resultData, nil
}
func (m *Module) GetEmployeeTotalRows(ctx context.Context, pg iModel.PaginationAndSearch) (int, error) {
	var (
		sb         strings.Builder
		resultData = 0
	)

	q := `select COUNT(employee_id)
		from public.employee`

	sb.WriteString(q)
	sb.WriteString(pg.BuildPaginationAndSearchQuery(false))

	err := m.sqlDbManager.Store().GetContext(ctx, &resultData, q)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, helper.NewSqlErr(err)
	}

	return resultData, nil
}

func (m *Module) GetEntityEmployeeByEmployeeId(ctx context.Context, employeeId uuid.UUID) (*entity.Employee, error) {
	var resultData = &entity.Employee{}

	q := fmt.Sprintf(`select
			employee_id,
			name,
			job_title,
			salary,
			department,
			joined_date
		from public.employee
		where employee_id=$1`)

	log.Info().Msg(q)

	err := m.sqlDbManager.Store().GetContext(ctx, resultData, q, employeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, helper.NewSqlErr(err)
	}

	return resultData, nil
}
