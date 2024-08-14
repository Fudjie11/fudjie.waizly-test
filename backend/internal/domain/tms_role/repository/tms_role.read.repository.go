package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	qGenRead "fudjie.waizly/backend-test/library/querygen/qgread"
	qGenReadModel "fudjie.waizly/backend-test/library/querygen/qgread/model"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	"fudjie.waizly/backend-test/internal/helper"
	tracer "fudjie.waizly/backend-test/library/tracer"
)

func (m *Module) GetTmsRoles(ctx context.Context, qop *model.TmsRoleQueryOp, includeDeleted bool) ([]model.TmsRoleDto, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.GetTmsRoles", nil)
	defer span.End()

	var (
		resultData = []model.TmsRoleDto{}
	)

	qDeleted := ""
	if !includeDeleted {
		qDeleted = "is_deleted=false"
	}

	q := `select 
			role_id,
			role_name,
			app_domain,
			description,
			is_deleted,
			is_active,
			created_user_id,
			created_time_utc,
			updated_user_id,
			updated_time_utc,
			row_version
		from public.tms_role
		$where
		$sort
		$paging`

	buildQueryParam := qGenReadModel.BuildQueryParam{
		Query:                   q,
		FilterSuffix:            qDeleted,
		DefaultSort:             "role_name asc",
		DefaultLimit:            100,
		StartingFilterArgNumber: 1,
	}

	qo := qGenReadModel.QueryOps[model.TmsRoleFilterQueryOp, model.TmsRoleSortQueryOp](*qop)
	qResult, qResultArgs, err := qGenRead.BuildFullQuery[model.TmsRoleFilterQueryOp, model.TmsRoleSortQueryOp](qo, buildQueryParam)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	log.Info().Msg(qResult)
	fmt.Println(qResultArgs)

	err = m.sqlDbManager.Store().SelectContext(ctx, &resultData, qResult, qResultArgs...)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, helper.NewSqlErr(err)
	}

	return resultData, nil
}

func (m *Module) GetTmsRoleTotalRows(ctx context.Context, qop *model.TmsRoleQueryOp, includeDeleted bool) (int, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.GetTmsRoleTotalRows", nil)
	defer span.End()

	qDeleted := ""
	if !includeDeleted {
		qDeleted = "is_deleted=false"
	}

	q := `select COUNT(role_id) 
		from public.tms_role
		$where`

	buildQueryParam := qGenReadModel.BuildQueryParam{
		Query:                   q,
		FilterSuffix:            qDeleted,
		DefaultSort:             "role_name asc",
		DefaultLimit:            100,
		StartingFilterArgNumber: 1,
	}

	qo := qGenReadModel.QueryOps[model.TmsRoleFilterQueryOp, model.TmsRoleSortQueryOp](*qop)
	qResult, qResultArgs, err := qGenRead.BuildFullQuery[model.TmsRoleFilterQueryOp, model.TmsRoleSortQueryOp](qo, buildQueryParam)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, err
	}

	log.Info().Msg(qResult)
	fmt.Println(qResultArgs)

	var totalRows int
	err = m.sqlDbManager.Store().GetContext(ctx, &totalRows, qResult, qResultArgs...)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, helper.NewSqlErr(err)
	}

	return totalRows, nil
}

func (m *Module) GetEntityTmsRoleByRoleId(ctx context.Context, roleId uuid.UUID, includeDeleted bool) (*entity.TmsRole, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.GetEntityTmsRoleByRoleId", nil)
	defer span.End()

	var resultData = &entity.TmsRole{}

	qDeleted := ""
	if !includeDeleted {
		qDeleted = "and is_deleted=false"
	}

	q := fmt.Sprintf(`select 
			role_id,
			role_name,
			app_domain,
			description,
			is_deleted,
			is_active,
			created_user_id,
			created_time_utc,
			updated_user_id,
			updated_time_utc,
			row_version
		from public.tms_role
		where role_id=$1 %s`, qDeleted)

	log.Info().Msg(q)
	fmt.Println(roleId)

	err := m.sqlDbManager.Store().GetContext(ctx, resultData, q, roleId)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return nil, helper.NewSqlErr(err)
	}

	return resultData, nil
}
