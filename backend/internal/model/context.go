package model

import "github.com/google/uuid"

type ContextAppData struct {
	AppId          string
	UserId         uuid.UUID
	TenantId       uuid.UUID
	CustomerId     uuid.UUID
	LanguageId     string
	TimeZoneId     string
	TimeZoneOffset int
}
