package handler

import (
	"fudjie.waizly/backend-test/internal/domain/employee/usecase"
	"github.com/gofiber/fiber/v2"
)

type EmployeeRestHandler interface {
	GetEmployees(ctx *fiber.Ctx) error
	GetDetailEmployeebyId(ctx *fiber.Ctx) error
	CreateEmployee(ctx *fiber.Ctx) error
	UpdateEmployee(ctx *fiber.Ctx) error
	DeleteEmployee(ctx *fiber.Ctx) error
}

type Module struct {
	employeeUseCase usecase.EmployeeUseCase
}

type Opts struct {
	EmployeeUseCase usecase.EmployeeUseCase
}

func New(o *Opts) EmployeeRestHandler {
	return &Module{
		employeeUseCase: o.EmployeeUseCase,
	}
}
