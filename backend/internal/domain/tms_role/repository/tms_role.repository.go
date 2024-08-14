package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	sqlStore "fudjie.waizly/backend-test/library/sqldb"
)

type TmsRoleRepository interface {
	GetTmsRoles(ctx context.Context, qop *model.TmsRoleQueryOp, includeDeleted bool) ([]model.TmsRoleDto, error)
	GetTmsRoleTotalRows(ctx context.Context, qop *model.TmsRoleQueryOp, includeDeleted bool) (int, error)
	GetEntityTmsRoleByRoleId(ctx context.Context, roleId uuid.UUID, includeDeleted bool) (*entity.TmsRole, error)

	CreateTmsRole(ctx context.Context, dbTx *sql.Tx, data *entity.TmsRole) error
	UpdateTmsRole(ctx context.Context, dbTx *sql.Tx, data *entity.TmsRole, rowVersion uuid.UUID) (int64, error)
	SoftDeleteTmsRole(ctx context.Context, dbTx *sql.Tx, roleId uuid.UUID, rowVersion uuid.UUID, actorUserId uuid.UUID) (int64, error)
}

type Module struct {
	sqlDbManager sqlStore.SqlDbManager
}

type Opts struct {
	SqlDbManager sqlStore.SqlDbManager
}

func New(o *Opts) TmsRoleRepository {
	return &Module{
		sqlDbManager: o.SqlDbManager,
	}
}
