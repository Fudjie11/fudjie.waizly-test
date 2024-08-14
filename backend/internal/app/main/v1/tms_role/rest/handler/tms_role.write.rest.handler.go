package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	pb "bluebird.tech/kirim/protobuf/auth/v1"
	structConverter "fudjie.waizly/backend-test/library/struct_converter"
	tracer "fudjie.waizly/backend-test/library/tracer"

	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	helper "fudjie.waizly/backend-test/internal/helper"
)

// @Summary Create Tms Role
// @Tags Tms - Role
// @Description api for create tms role
// @Accept json
// @Produce json
// @Param payload body pb.CreateTmsRoleRequest true "payload"
// @Success 200 {object} pb.CreateTmsRoleResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/tms-role/create [POST]
func (m *Module) CreateTmsRole(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	_, span := tracer.StartSpan(ctx, "rest.handler.tms_role.CreateTmsRole", nil)
	defer span.End()

	var (
		req      = &pb.CreateTmsRoleRequest{}
		dto      = &model.AddTmsRoleDto{}
		response = &pb.CreateTmsRoleResponse{}
		err      error
	)

	if err = c.BodyParser(&req); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto, err = structConverter.ConvertStruct[*model.AddTmsRoleDto](req)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto.PermissionIds = req.GetPermissionIds()

	dataId, err := m.tmsRoleUseCase.CreateTmsRole(ctx, dto)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	response.Data = &pb.DataIdResponse{
		DataId: dataId.String(),
	}

	return c.JSON(response)
}

// @Summary Update Tms Role
// @Tags Tms - Role
// @Description api for update tms role by id
// @Accept json
// @Produce json
// @Param tmsRoleId path string true "Tms Role Id" format(uuid)
// @Param payload body pb.UpdateTmsRoleRequest true "Update data for tms role"
// @Success 200 {object} pb.DetailTmsRole
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/tms-role/{tmsRoleId}/update [PUT]
func (m *Module) UpdateTmsRole(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	_, span := tracer.StartSpan(ctx, "rest.handler.tms_role.UpdateTmsRole", nil)
	defer span.End()

	var (
		req      = &pb.UpdateTmsRoleRequest{}
		dto      = &model.UpdateTmsRoleDto{}
		response = &pb.UpdateTmsRoleResponse{}
		err      error
	)

	if err = c.BodyParser(&req); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	tmsRoleId, err := uuid.Parse(c.Params("tmsRoleId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto, err = structConverter.ConvertStruct[*model.UpdateTmsRoleDto](req)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto.PermissionIds = req.GetPermissionIds()
	dto.RoleId = tmsRoleId

	err = m.tmsRoleUseCase.UpdateTmsRole(ctx, dto)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	response.Data = &pb.DataIdResponse{
		DataId: dto.RoleId.String(),
	}

	return c.JSON(response)
}

// @Summary Delete Tms Role
// @Tags Tms - Role
// @Description api for delete Tms Role by Id
// @Accept json
// @Produce json
// @Param tmsRoleId path string true "Tms Role Id" format(uuid)
// @Param payload body pb.DeleteTmsRoleRequest true "delete data for tms role"
// @Success 200 {object} pb.DeleteTmsRoleResponse
// @Failure 400 {object} pb.DeleteTmsRoleResponse
// @Failure 500 {object} pb.DeleteTmsRoleResponse
// @Security Bearer
// @Router /api/main/v1/tms-role/{tmsRoleId}/delete [DELETE]
func (m *Module) DeleteTmsRole(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	_, span := tracer.StartSpan(ctx, "rest.handler.tms_role.UpdateTmsRole", nil)
	defer span.End()

	var (
		req      = &pb.DeleteTmsRoleRequest{}
		dto      = &model.DeleteTmsRoleDto{}
		response = &pb.DeleteTmsRoleResponse{}
		err      error
	)

	if err = c.BodyParser(&req); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	tmsRoleId, err := uuid.Parse(c.Params("tmsRoleId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto, err = structConverter.ConvertStruct[*model.DeleteTmsRoleDto](req)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	dto.RoleId = tmsRoleId

	err = m.tmsRoleUseCase.SoftDeleteTmsRole(ctx, dto)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	response.Data = &pb.DataIdResponse{
		DataId: dto.RoleId.String(),
	}

	return c.JSON(response)
}
