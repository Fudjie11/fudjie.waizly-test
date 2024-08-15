package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	entity "fudjie.waizly/backend-test/internal/db/entity"
	model "fudjie.waizly/backend-test/internal/domain/employee/model"
	"fudjie.waizly/backend-test/internal/helper"
)

func (m *Module) CreateEmployee(ctx context.Context, dto *model.AddEmployeeDto) (uuid.UUID, error) {
	var (
		err error
	)

	eEmployee := entity.NewEmployee()
	eEmployee.EmployeeId = uuid.New()
	eEmployee.Name = dto.Name
	eEmployee.JobTitle = dto.JobTitle
	eEmployee.Salary = dto.Salary
	eEmployee.Department = dto.Department
	eEmployee.JoinedDate = time.Now()

	if err = m.employeeRepository.CreateEmployee(ctx, eEmployee); err != nil {
		log.Err(err).Msg(err.Error())
		return uuid.Nil, err
	}

	return eEmployee.EmployeeId, err
}

func (m *Module) UpdateEmployee(ctx context.Context, dto *model.UpdateEmployeeDto) error {
	var (
		eEmployee = &entity.Employee{}
		err       error
	)

	eEmployee, err = m.employeeRepository.GetEntityEmployeeByEmployeeId(ctx, dto.EmployeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	if eEmployee == nil {
		return helper.NewNotFoundErr(fmt.Errorf("employee not found"))
	}

	eEmployee.Name = dto.Name
	eEmployee.JobTitle = dto.JobTitle
	eEmployee.Salary = dto.Salary
	eEmployee.Department = dto.Department
	eEmployee.JoinedDate = time.Now()

	_, err = m.employeeRepository.UpdateEmployee(ctx, eEmployee)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (m *Module) DeleteEmployee(ctx context.Context, employeeId uuid.UUID) error {
	eEmployee, err := m.employeeRepository.GetEntityEmployeeByEmployeeId(ctx, employeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	if eEmployee == nil {
		return helper.NewNotFoundErr(fmt.Errorf("employee not found"))
	}

	_, err = m.employeeRepository.DeleteEmployee(ctx, employeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}
