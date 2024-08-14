package model

import (
	"time"

	"github.com/google/uuid"

	structConverter "fudjie.waizly/backend-test/library/struct_converter"

	entity "fudjie.waizly/backend-test/internal/db/auth/entity"
)

type TmsRoleDto struct {
	RoleId         uuid.UUID `json:"role_id" db:"role_id" validate:"required"`
	RoleName       string    `json:"role_name" db:"role_name" validate:"required"`
	AppDomain      string    `json:"app_domain" db:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	Description    string    `json:"description" db:"description"`
	IsActive       bool      `json:"is_active" db:"is_active" validate:"required"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted" validate:"required"`
	CreatedUserId  uuid.UUID `json:"created_user_id" db:"created_user_id" validate:"required"`
	CreatedTimeUtc time.Time `json:"created_time_utc" db:"created_time_utc" validate:"required"`
	UpdatedUserId  uuid.UUID `json:"updated_user_id" db:"updated_user_id" validate:"required"`
	UpdatedTimeUtc time.Time `json:"updated_time_utc" db:"updated_time_utc" validate:"required"`
	RowVersion     uuid.UUID `json:"row_version" db:"row_version" validate:"required"`
}

type DetailTmsRoleDto struct {
	RoleId         uuid.UUID `json:"role_id" db:"role_id" validate:"required"`
	RoleName       string    `json:"role_name" db:"role_name" validate:"required"`
	AppDomain      string    `json:"app_domain" db:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	Description    string    `json:"description" db:"description"`
	IsActive       bool      `json:"is_active" db:"is_active" validate:"required"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted" validate:"required"`
	CreatedUserId  uuid.UUID `json:"created_user_id" db:"created_user_id" validate:"required"`
	CreatedTimeUtc time.Time `json:"created_time_utc" db:"created_time_utc" validate:"required"`
	UpdatedUserId  uuid.UUID `json:"updated_user_id" db:"updated_user_id" validate:"required"`
	UpdatedTimeUtc time.Time `json:"updated_time_utc" db:"updated_time_utc" validate:"required"`
	RowVersion     uuid.UUID `json:"row_version" db:"row_version" validate:"required"`

	// Relation
	Permission []*Permission `json:"permission" db:"permission"`
}

type Permission struct {
	PermissionId       string `json:"permission_id" validate:"required,max=50"`
	AppDomain          string `json:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	PermissionName     string `json:"permission_name" validate:"required,max=100" transform:"trim"`
	PermissionGroup    string `json:"permission_group" validate:"max=100" transform:"trim"`
	PermissionSubGroup string `json:"permission_sub_group" validate:"max=100" transform:"trim"`
	Description        string `json:"description"`
}

type RolePermissionQeuryResult struct {
	RoleId             uuid.UUID `json:"role_id" db:"role_id"`
	PermissionId       string    `json:"permission_id"`
	AppDomain          string    `json:"app_domain"`
	PermissionName     string    `json:"permission_name"`
	PermissionGroup    string    `json:"permission_group"`
	PermissionSubGroup string    `json:"permission_sub_group"`
	Description        string    `json:"description"`
}

type AddTmsRoleDto struct {
	RoleName      string    `json:"role_name" validate:"required,max=50" transform:"trim"`
	AppDomain     string    `json:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	PermissionIds []string  `json:"permissions_ids"`
	Description   string    `json:"description"`
	IsActive      bool      `json:"is_active"`
	ActorUserId   uuid.UUID `json:"actor_user_id" validate:"required"`
}

type UpdateTmsRoleDto struct {
	RoleId        uuid.UUID `json:"role_id" validate:"required"`
	RoleName      string    `json:"role_name" validate:"required,max=50" transform:"trim"`
	AppDomain     string    `json:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	PermissionIds []string  `json:"permissions_ids" json:"permissions"`
	Description   string    `json:"description"`
	IsActive      bool      `json:"is_active"`
	ActorUserId   uuid.UUID `json:"actor_user_id" validate:"required"`
	RowVersion    uuid.UUID `json:"row_version" validate:"required"`
}

type DeleteTmsRoleDto struct {
	RoleId      uuid.UUID `json:"role_id" validate:"required"`
	RowVersion  uuid.UUID `json:"row_version" validate:"required"`
	ActorUserId uuid.UUID `json:"actor_user_id" validate:"required"`
}

func ToPermissionsDto(ePermission []entity.TmsPermission) []*Permission {
	var permissions []*Permission

	for _, v := range ePermission {
		permission := structConverter.MustConvertStruct[Permission](v)
		permissions = append(permissions, &permission)
	}

	return permissions
}
