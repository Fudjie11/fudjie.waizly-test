package router

import (
	"fudjie.waizly/backend-test/internal/app/main/v1/employee/rest/handler"
	"fudjie.waizly/backend-test/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Configure(app fiber.Router, middleware middleware.RestMiddleware, employeeHandler handler.EmployeeRestHandler) {
	app.Post(
		"/employees",
		middleware.GuardBasicAuthentication,
		employeeHandler.GetEmployees,
	)

	app.Get(
		"/employee/:employeeId",
		middleware.GuardBasicAuthentication,
		employeeHandler.GetDetailEmployeebyId,
	)

	app.Post(
		"/employee/create",
		middleware.GuardBasicAuthentication,
		employeeHandler.CreateEmployee,
	)

	app.Put(
		"/employee/:employeeId/update",
		middleware.GuardBasicAuthentication,
		employeeHandler.UpdateEmployee,
	)

	app.Delete(
		"/employee/:employeeId/delete",
		middleware.GuardBasicAuthentication,
		employeeHandler.DeleteEmployee,
	)

}
