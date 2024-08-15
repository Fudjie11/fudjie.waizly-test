package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/rs/zerolog/log"

	model "fudjie.waizly/backend-test/internal/domain/employee/model"
	helper "fudjie.waizly/backend-test/internal/helper"
	iModel "fudjie.waizly/backend-test/internal/model"
)

// @Summary Get List Employee
// @Tags Employee
// @Description api for Get list tms role
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param payload body iModel.PaginationAndSearch true "payload"
// @Success 200 {object} iModel.ListPaging[model.EmployeeDto]
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/employees [POST]
func (m *Module) GetEmployees(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	var (
		pg  iModel.PaginationAndSearch
		err error
	)

	if err = c.BodyParser(&pg); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	results, err := m.employeeUseCase.GetEmployees(ctx, pg)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return helper.NewOkPagingResponse(c, results)
}

// @Summary Get Employee  By Id
// @Tags Employee
// @Description api for Get Employee  by Id
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param employeeId path string true "Employee Id" format(uuid)
// @Success 200 {object} model.EmployeeDto
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/employee/{employeeId} [GET]
func (m *Module) GetDetailEmployeebyId(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)

	var (
		detailemployee = &model.EmployeeDto{}
	)

	employeeId, err := uuid.Parse(c.Params("employeeId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	result, err := m.employeeUseCase.GetEmployeeByEmployeeId(ctx, employeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	detailemployee.EmployeeId = result.EmployeeId
	detailemployee.Name = result.Name
	detailemployee.JobTitle = result.JobTitle
	detailemployee.Salary = result.Salary
	detailemployee.Department = result.Department
	detailemployee.JoinedDate = result.JoinedDate

	return c.JSON(result)
}
