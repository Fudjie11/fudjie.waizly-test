package handler

import (
	"github.com/gofiber/fiber/v2"

	"fudjie.waizly/backend-test/internal/domain/tms_role/usecase"
)

type TmsRoleRestHandler interface {
	GetTmsRoles(ctx *fiber.Ctx) error
	GetDetailTmsRolesbyId(ctx *fiber.Ctx) error
	CreateTmsRole(ctx *fiber.Ctx) error
	UpdateTmsRole(ctx *fiber.Ctx) error
	DeleteTmsRole(ctx *fiber.Ctx) error
}

type Module struct {
	tmsRoleUseCase usecase.TmsRoleUseCase
}

type Opts struct {
	TmsRoleUseCase usecase.TmsRoleUseCase
}

func New(o *Opts) TmsRoleRestHandler {
	return &Module{
		tmsRoleUseCase: o.TmsRoleUseCase,
	}
}
