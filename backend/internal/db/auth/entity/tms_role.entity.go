package entity

import (
	"context"
	"time"

	"github.com/google/uuid"

	ienum "fudjie.waizly/backend-test/internal/enum"
	ihelper "fudjie.waizly/backend-test/internal/helper"
)

type TmsRole struct {
	RoleId         uuid.UUID `bson:"role_id" json:"role_id" db:"role_id" validate:"required"`
	RoleName       string    `bson:"role_name" json:"role_name" db:"role_name" validate:"required"`
	AppDomain      string    `bson:"app_domain" json:"app_domain" db:"app_domain" validate:"required,max=30" transform:"lower,trim"`
	Description    string    `bson:"description" json:"description" db:"description" validate:"required"`
	IsActive       bool      `bson:"is_active" json:" is_active" db:"is_active"`
	IsDeleted      bool      `bson:"is_deleted" json:"is_deleted" db:"is_deleted"`
	CreatedUserId  uuid.UUID `bson:"created_user_id" json:"created_user_id" db:"created_user_id" validate:"required"`
	CreatedTimeUtc time.Time `bson:"created_time_utc" json:"created_time_utc" db:"created_time_utc" validate:"required"`
	UpdatedUserId  uuid.UUID `bson:"updated_user_id" json:"updated_user_id" db:"updated_user_id" validate:"required"`
	UpdatedTimeUtc time.Time `bson:"updated_time_utc" json:"updated_time_utc" db:"updated_time_utc" validate:"required"`
	RowVersion     uuid.UUID `bson:"row_version" json:"row_version" db:"row_version" validate:"required"`
}

func NewTmsRole() *TmsRole {
	instance := &TmsRole{
		IsDeleted: false,
		IsActive:  true,
	}

	return instance
}

func (r *TmsRole) FinalizeAndValidate(ctx context.Context, actionType ienum.CRUDAction, actorUserId uuid.UUID) error {
	utcTime := time.Now().UTC()

	if actionType == ienum.CRUDActionInsert {
		r.CreatedUserId = actorUserId
		r.CreatedTimeUtc = utcTime
	}

	r.UpdatedUserId = actorUserId
	r.UpdatedTimeUtc = utcTime
	r.RowVersion = uuid.New()

	if err := ihelper.TransformAndValidateByTag(ctx, r); err != nil {
		return err
	}

	return nil
}
