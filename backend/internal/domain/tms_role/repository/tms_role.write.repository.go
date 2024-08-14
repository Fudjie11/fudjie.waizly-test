package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	"fudjie.waizly/backend-test/internal/helper"
	qGenWrite "fudjie.waizly/backend-test/library/querygen/qgwrite"
	tracer "fudjie.waizly/backend-test/library/tracer"
)

func (m *Module) CreateTmsRole(ctx context.Context, dbTx *sql.Tx, data *entity.TmsRole) error {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.CreateTmsRole", nil)
	defer span.End()

	qResult, qResultArgs, err := qGenWrite.GenerateInsertQueryFromEntity(data, "public.tms_role", nil, nil)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	log.Info().Msg(qResult)
	fmt.Println(qResultArgs)

	_, err = dbTx.ExecContext(ctx, qResult, qResultArgs...)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return helper.NewSqlErr(err)
	}

	return nil
}

func (m *Module) UpdateTmsRole(ctx context.Context, dbTx *sql.Tx, data *entity.TmsRole, rowVersion uuid.UUID) (int64, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.UpdateTmsRole", nil)
	defer span.End()

	qResult, qResultArgs, err := qGenWrite.GenerateUpdateQueryFromEntityForAllFields(
		data,
		"public.tms_role",
		"role_id=$1 and row_version=$2",
		3,
		[]interface{}{data.RoleId, rowVersion},
		[]string{"role_id", "is_deleted", "created_user_id", "created_time_utc"},
		nil,
	)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, err
	}

	log.Info().Msg(qResult)
	fmt.Println(qResultArgs)

	res, err := dbTx.ExecContext(ctx, qResult, qResultArgs...)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, helper.NewSqlErr(err)
	}

	return res.RowsAffected()
}

func (m *Module) SoftDeleteTmsRole(ctx context.Context, dbTx *sql.Tx, roleId uuid.UUID, rowVersion uuid.UUID, actorUserId uuid.UUID) (rowsAffected int64, err error) {
	ctx, span := tracer.StartSpan(ctx, "repo.tms_role.SoftDeleteTmsRole", nil)
	defer span.End()

	updatedTime := time.Now().UTC()
	newRowVersion := uuid.New()

	res, err := dbTx.ExecContext(
		ctx,
		`update public.tms_role 
		set is_deleted=true, updated_user_id=$3, updated_time_utc=$4, row_version=$5
		where role_id=$1 and row_version=$2`,
		roleId,
		rowVersion,
		actorUserId,
		updatedTime,
		newRowVersion,
	)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return 0, helper.NewSqlErr(err)
	}

	return res.RowsAffected()
}
