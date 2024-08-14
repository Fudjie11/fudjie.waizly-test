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

// @Summary Get List Tms Role
// @Tags Tms - Role
// @Description api for Get list tms role
// @Accept json
// @Produce json
// @Param payload body pb.GetTmsRolesRequest true "payload"
// @Success 200 {object} pb.GetTmsRolesResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/tms-roles [POST]
func (m *Module) GetTmsRoles(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	_, span := tracer.StartSpan(ctx, "rest.handler.tms_role.GetTmsRoles", nil)
	defer span.End()
	var (
		qop model.TmsRoleQueryOp
		err error
	)

	if qop, err = helper.ParsePagingFromHttp[model.TmsRoleFilterQueryOp, model.TmsRoleSortQueryOp](c, qop); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	results, err := m.tmsRoleUseCase.GetTmsRoles(ctx, &qop)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return helper.NewOkPagingResponse(c, results)
}

// @Summary Get Tms Role  By Id
// @Tags Tms - Role
// @Description api for Get Tms Role  by Id
// @Accept json
// @Produce json
// @Param tmsRoleId path string true "Tms Role Id" format(uuid)
// @Success 200 {object} pb.GetDetailTmsRolesbyIdResponse
// @Failure 400 {object} helper.JSONResponse
// @Failure 500 {object} helper.JSONResponse
// @Security Bearer
// @Router /api/main/v1/tms-role/{tmsRoleId} [GET]
func (m *Module) GetDetailTmsRolesbyId(c *fiber.Ctx) error {
	ctx := helper.NewContextFromHttp(c)
	_, span := tracer.StartSpan(ctx, "rest.handler.tms_role.GetTmsRoles", nil)
	defer span.End()
	var (
		pbDetailTmsRole = &pb.DetailTmsRole{}
		response        = &pb.GetDetailTmsRolesbyIdResponse{}
	)

	tmsRoleId, err := uuid.Parse(c.Params("tmsRoleId"))
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	result, err := m.tmsRoleUseCase.GetTmsRoleByRoleId(ctx, tmsRoleId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	pbDetailTmsRole, err = structConverter.ConvertStruct[*pb.DetailTmsRole](result)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewInternalServerErr(err)
	}

	response.Data = pbDetailTmsRole

	return c.JSON(response)
}
