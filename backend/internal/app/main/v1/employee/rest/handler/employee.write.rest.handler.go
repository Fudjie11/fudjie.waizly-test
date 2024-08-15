package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	model "fudjie.waizly/backend-test/internal/domain/employee/model"
	helper "fudjie.waizly/backend-test/internal/helper"
	iModel "fudjie.waizly/backend-test/internal/model"
)

// @Summary Create Employee
// @Tags Employee
// @Description api for create tms role
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param payload body model.AddEmployeeDto true "payload"
// @Success 200 {object} iModel.WriteResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/employee/create [POST]
func (m *Module) CreateEmployee(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)

	var (
		req = &model.AddEmployeeDto{}
		res = &iModel.WriteResponse{}
		err error
	)

	if err = c.BodyParser(&req); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dataId, err := m.employeeUseCase.CreateEmployee(ctx, req)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	res.DataId = dataId
	res.Success = true

	return c.JSON(res)
}

// @Summary Update Employee
// @Tags Employee
// @Description api for update tms role by id
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param employeeId path string true "Employee Id" format(uuid)
// @Param payload body model.UpdateEmployeeDto true "Update data for tms role"
// @Success 200 {object} iModel.WriteResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/employee/{employeeId}/update [PUT]
func (m *Module) UpdateEmployee(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)

	var (
		req = &model.UpdateEmployeeDto{}
		res = &iModel.WriteResponse{}
		err error
	)

	if err = c.BodyParser(&req); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	employeeId, err := uuid.Parse(c.Params("employeeId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	req.EmployeeId = employeeId

	err = m.employeeUseCase.UpdateEmployee(ctx, req)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	res.DataId = employeeId
	res.Success = true

	return c.JSON(res)
}

// @Summary Delete Employee
// @Tags Employee
// @Description api for delete Employee by Id
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param employeeId path string true "Employee Id" format(uuid)
// @Success 200 {object} iModel.WriteResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/employee/{employeeId}/delete [DELETE]
func (m *Module) DeleteEmployee(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)

	var (
		err error
		res = &iModel.WriteResponse{}
	)

	employeeId, err := uuid.Parse(c.Params("employeeId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	err = m.employeeUseCase.DeleteEmployee(ctx, employeeId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	res.DataId = employeeId
	res.Success = true

	return c.JSON(res)
}
