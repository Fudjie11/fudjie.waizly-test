package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	structConverter "fudjie.waizly/backend-test/library/struct_converter"
	tracer "fudjie.waizly/backend-test/library/tracer"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	shModel "fudjie.waizly/backend-test/internal/model"
)

func (m *Module) GetTmsRoles(ctx context.Context, qop *model.TmsRoleQueryOp) (*shModel.ListPaging[model.TmsRoleDto], error) {
	ctx, span := tracer.StartSpan(ctx, "usecase.tms_role.GetTmsRoles", nil)
	defer span.End()

	var (
		err       error
		results   []model.TmsRoleDto
		totalRows int
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		results, err = m.tmsRoleRepository.GetTmsRoles(gctx, qop, false)
		return err
	})

	g.Go(func() error {
		totalRows, err = m.tmsRoleRepository.GetTmsRoleTotalRows(gctx, qop, false)
		return err
	})

	if err = g.Wait(); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	resultData := shModel.NewListPaging(results, totalRows, qop.Limit, qop.Offset, qop.Page)

	return &resultData, nil
}

func (m *Module) GetTmsRoleByRoleId(ctx context.Context, roleId uuid.UUID) (*model.DetailTmsRoleDto, error) {
	ctx, span := tracer.StartSpan(ctx, "usecase.tms_role.GetTmsRoleByRoleId", nil)
	defer span.End()

	var (
		eTmsRole    *entity.TmsRole
		ePermission []entity.TmsPermission
		err         error
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		eTmsRole, err = m.tmsRoleRepository.GetEntityTmsRoleByRoleId(gctx, roleId, false)
		return err
	})

	g.Go(func() error {
		ePermission, err = m.tmsRoleToPermissionRepository.GetEntityPermissionByRoleId(gctx, roleId, false)
		return err
	})

	if err = g.Wait(); err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	detailTmsRoleDto := structConverter.MustConvertStruct[model.DetailTmsRoleDto](eTmsRole)
	tmsPermissionDto := model.ToPermissionsDto(ePermission)

	detailTmsRoleDto.Permission = tmsPermissionDto

	return &detailTmsRoleDto, nil
}
