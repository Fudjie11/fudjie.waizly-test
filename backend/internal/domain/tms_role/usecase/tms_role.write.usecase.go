package usecase

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	iEnum "fudjie.waizly/backend-test/internal/enum"
	helper "fudjie.waizly/backend-test/internal/helper"
	tracer "fudjie.waizly/backend-test/library/tracer"
)

func (m *Module) CreateTmsRole(ctx context.Context, dto *model.AddTmsRoleDto) (uuid.UUID, error) {
	ctx, span := tracer.StartSpan(ctx, "usecase.tms_role.CreateTmsRole", nil)
	defer span.End()

	var (
		err                   error
		eTmsRoleToPermissions []entity.TmsRoleToPermission
	)

	if err := helper.TransformAndValidateByTag(ctx, dto); err != nil {
		log.Err(err).Msg(err.Error())
		return uuid.Nil, helper.NewBadRequestErr(err)
	}

	eTmsRole := entity.NewTmsRole()
	eTmsRole.RoleId = uuid.New()
	eTmsRole.RoleName = dto.RoleName
	eTmsRole.AppDomain = dto.AppDomain
	eTmsRole.Description = dto.RoleName

	if err := eTmsRole.FinalizeAndValidate(ctx, iEnum.CRUDActionInsert, dto.ActorUserId); err != nil {
		log.Err(err).Msg(err.Error())
		return uuid.Nil, helper.NewBadRequestErr(err)
	}

	for _, permission := range dto.PermissionIds {
		eTmsRoleToPermission := entity.NewTmsRolePermission()
		eTmsRoleToPermission.RoleId = eTmsRole.RoleId
		eTmsRoleToPermission.PermissionId = permission

		if err := eTmsRoleToPermission.FinalizeAndValidate(ctx, iEnum.CRUDActionInsert, dto.ActorUserId); err != nil {
			log.Err(err).Msg(err.Error())
			return uuid.Nil, helper.NewBadRequestErr(err)
		}

		eTmsRoleToPermissions = append(eTmsRoleToPermissions, *eTmsRoleToPermission)
	}

	err = m.sqlDbManager.WrapTransaction(ctx, func(ctx context.Context, dbTx *sql.Tx) error {
		// Create TMs Role
		if err = m.tmsRoleRepository.CreateTmsRole(ctx, dbTx, eTmsRole); err != nil {
			return err
		}

		// Create bulk Tms Role To Permission
		if err = m.tmsRoleToPermissionRepository.CreateBulkTmsRoleToPermission(ctx, dbTx, eTmsRoleToPermissions); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Err(err).Msg(err.Error())
		return uuid.Nil, err
	}

	return eTmsRole.RoleId, err
}

func (m *Module) UpdateTmsRole(ctx context.Context, dto *model.UpdateTmsRoleDto) error {
	ctx, span := tracer.StartSpan(ctx, "usecase.tms_role.UpdateTmsRole", nil)
	defer span.End()
	var (
		eTmsRole                 = &entity.TmsRole{}
		eTmsPermissions          []entity.TmsPermission
		newETmsRoleToPermissions []entity.TmsRoleToPermission
		err                      error
	)

	if err := helper.TransformAndValidateByTag(ctx, dto); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewBadRequestErr(err)
	}

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		eTmsRole, err = m.tmsRoleRepository.GetEntityTmsRoleByRoleId(gctx, dto.RoleId, false)
		if err != nil {
			log.Err(err).Msg(err.Error())
		}
		return err
	})

	g.Go(func() error {
		eTmsPermissions, err = m.tmsRoleToPermissionRepository.GetEntityPermissionByRoleId(gctx, dto.RoleId, false)
		if err != nil {
			log.Err(err).Msg(err.Error())
		}
		return err
	})

	if err = g.Wait(); err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	eTmsRole.RoleName = dto.RoleName
	eTmsRole.AppDomain = dto.AppDomain
	eTmsRole.Description = dto.RoleName
	eTmsRole.IsActive = dto.IsActive
	oriTmsRoleRowVersion := dto.RowVersion

	if err := eTmsRole.FinalizeAndValidate(ctx, iEnum.CRUDActionUpdate, dto.ActorUserId); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewBadRequestErr(err)
	}

	newPermissionIds, deletePermissionIds := m.getNewAndDeletPermission(dto.PermissionIds, eTmsPermissions)

	for _, permissionId := range newPermissionIds {
		eTmsRoleToPermission := entity.NewTmsRolePermission()
		eTmsRoleToPermission.RoleId = dto.RoleId
		eTmsRoleToPermission.PermissionId = permissionId

		if err := eTmsRoleToPermission.FinalizeAndValidate(ctx, iEnum.CRUDActionInsert, dto.ActorUserId); err != nil {
			log.Err(err).Msg(err.Error())
			return helper.NewBadRequestErr(err)
		}

		newETmsRoleToPermissions = append(newETmsRoleToPermissions, *eTmsRoleToPermission)
	}

	err = m.sqlDbManager.WrapTransaction(ctx, func(ctx context.Context, dbTx *sql.Tx) error {
		// Update Tms Role
		if err = helper.RowAffectedCheckerWithoutLocalizer(func() (int64, error) {
			return m.tmsRoleRepository.UpdateTmsRole(ctx, dbTx, eTmsRole, oriTmsRoleRowVersion)
		}); err != nil {
			return err
		}

		if len(newETmsRoleToPermissions) != 0 {
			// Create New Relation role to permission
			if err = m.tmsRoleToPermissionRepository.CreateBulkTmsRoleToPermission(ctx, dbTx, newETmsRoleToPermissions); err != nil {
				return err
			}
		}

		if len(deletePermissionIds) != 0 {
			// Delete Bulk Tms Role To Permission by PermissionIds
			if err = m.tmsRoleToPermissionRepository.DeleteBulkTmsRoleToPermissionbyPermissionIds(ctx, dbTx, deletePermissionIds, eTmsRole.RoleId); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}

func (m *Module) SoftDeleteTmsRole(ctx context.Context, dto *model.DeleteTmsRoleDto) error {
	ctx, span := tracer.StartSpan(ctx, "usecase.tms_role.SoftDeleteTmsRole", nil)
	defer span.End()

	if err := helper.TransformAndValidateByTag(ctx, dto); err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewBadRequestErr(err)
	}

	eTmsRole, err := m.tmsRoleRepository.GetEntityTmsRoleByRoleId(ctx, dto.RoleId, false)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	err = m.sqlDbManager.WrapTransaction(ctx, func(ctx context.Context, dbTx *sql.Tx) error {
		// Update Tms Role
		if err = helper.RowAffectedCheckerWithoutLocalizer(func() (int64, error) {
			return m.tmsRoleRepository.SoftDeleteTmsRole(ctx, dbTx, eTmsRole.RoleId, dto.RowVersion, dto.ActorUserId)
		}); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}
