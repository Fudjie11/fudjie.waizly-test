package usecase

import (
	"context"

	"github.com/google/uuid"

	model "fudjie.waizly/backend-test/internal/domain/tms_role/model"
	tmsRoleRepo "fudjie.waizly/backend-test/internal/domain/tms_role/repository"

	// keycloakApi "fudjie.waizly/backend-test/internal/integration/keycloak/api"
	shModel "fudjie.waizly/backend-test/internal/model"
	sqlStore "fudjie.waizly/backend-test/library/sqldb"
)

type TmsRoleUseCase interface {
	GetTmsRoles(ctx context.Context, qop *model.TmsRoleQueryOp) (*shModel.ListPaging[model.TmsRoleDto], error)
	GetTmsRoleByRoleId(ctx context.Context, branchId uuid.UUID) (*model.DetailTmsRoleDto, error)

	CreateTmsRole(ctx context.Context, dto *model.AddTmsRoleDto) (uuid.UUID, error)
	UpdateTmsRole(ctx context.Context, dto *model.UpdateTmsRoleDto) error
	SoftDeleteTmsRole(ctx context.Context, dto *model.DeleteTmsRoleDto) error
}

type Module struct {
	sqlDbManager                  sqlStore.SqlDbManager
	tmsRoleRepository             tmsRoleRepo.TmsRoleRepository
	tmsRoleToPermissionRepository tmsRoleRepo.TmsRoleToPermissionRepository
	// keycloakApiRealmInternalEmployee keycloakApi.KeycloakApi
}

type Opts struct {
	SqlDbManager                  sqlStore.SqlDbManager
	TmsRoleRepository             tmsRoleRepo.TmsRoleRepository
	TmsRoleToPermissionRepository tmsRoleRepo.TmsRoleToPermissionRepository
	// KeycloakApiRealmInternalEmployee keycloakApi.KeycloakApi
}

func New(o *Opts) TmsRoleUseCase {
	return &Module{
		sqlDbManager:                  o.SqlDbManager,
		tmsRoleRepository:             o.TmsRoleRepository,
		tmsRoleToPermissionRepository: o.TmsRoleToPermissionRepository,
		// keycloakApiRealmInternalEmployee: o.KeycloakApiRealmInternalEmployee,
	}
}
