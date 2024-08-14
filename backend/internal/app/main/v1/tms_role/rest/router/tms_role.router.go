package router

import (
	"fudjie.waizly/backend-test/library/internal/app/main/v1/tms_role/rest/handler"
	"fudjie.waizly/backend-test/library/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Configure(app fiber.Router, middleware middleware.HttpMiddleware, tmsRoleHandler handler.TmsRoleRestHandler) {
	app.Post(
		"/tms-roles",
		middleware.RequestDataMiddleware,
		middleware.ResponseDataMiddleware,
		tmsRoleHandler.GetTmsRoles,
	)

	app.Get(
		"/tms-role/:tmsRoleId",
		middleware.RequestDataMiddleware,
		middleware.ResponseDataMiddleware,
		tmsRoleHandler.GetDetailTmsRolesbyId,
	)

	app.Post(
		"/tms-role/create",
		middleware.RequestDataMiddleware,
		middleware.ResponseDataMiddleware,
		tmsRoleHandler.CreateTmsRole,
	)

	app.Put(
		"/tms-role/:tmsRoleId/update",
		middleware.RequestDataMiddleware,
		middleware.ResponseDataMiddleware,
		tmsRoleHandler.UpdateTmsRole,
	)

	app.Delete(
		"/tms-role/:tmsRoleId/delete",
		middleware.RequestDataMiddleware,
		middleware.ResponseDataMiddleware,
		tmsRoleHandler.DeleteTmsRole,
	)

}
