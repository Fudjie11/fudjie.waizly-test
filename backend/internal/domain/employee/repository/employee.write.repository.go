package repository

import (
	"context"

	"fudjie.waizly/backend-test/internal/db/entity"
	"fudjie.waizly/backend-test/internal/helper"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (m *Module) CreateEmployee(ctx context.Context, data *entity.Employee) error {
	queryInsertEmployee := `
	INSERT INTO employee (
		employee_id,
		name,
		job_title,
		salary,
		department,
		joined_date
	)
	VALUES (?,?,?,?,?,?);`

	q := m.sqlDbManager.Store().Rebind(queryInsertEmployee)
	_, err := m.sqlDbManager.Store().ExecContext(ctx, q,
		data.EmployeeId,
		data.Name,
		data.JobTitle,
		data.Salary,
		data.Department,
		data.JoinedDate,
	)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewSqlErr(err)
	}

	return nil
}

func (m *Module) UpdateEmployee(ctx context.Context, data *entity.Employee) (int64, error) {
	queryUpdateEmployee := `
		UPDATE employee 
		SET 
			name = ?,
			job_title = ?,
			salary = ?,
			department = ?,
			joined_date = ?,
		WHERE 
			employee_id = ?;
	`

	q := m.sqlDbManager.Store().Rebind(queryUpdateEmployee)
	log.Info().Msg(q)

	res, err := m.sqlDbManager.Store().ExecContext(ctx, q,
		data.Name,
		data.JobTitle,
		data.Salary,
		data.Department,
		data.JoinedDate,
		data.EmployeeId,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (m *Module) DeleteEmployee(ctx context.Context, employeeId uuid.UUID) (rowsAffected int64, err error) {
	queryDeleteEmployee := `
		DELETE FROM employee WHERE employee_id = ?;`

	q := m.sqlDbManager.Store().Rebind(queryDeleteEmployee)
	log.Info().Msg(q)

	res, err := m.sqlDbManager.Store().ExecContext(ctx, q,
		employeeId,
	)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, helper.NewSqlErr(err)
	}

	return res.RowsAffected()
}
