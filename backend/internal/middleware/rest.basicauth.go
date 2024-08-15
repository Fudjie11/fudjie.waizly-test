package middleware

import (
	"fudjie.waizly/backend-test/config"
	"github.com/gofiber/fiber/v2"
)

type RestMiddleware interface {
	GuardBasicAuthentication(fc *fiber.Ctx) error
}

type RestMiddlewareModule struct {
	authConfig config.AuthorizationConfig
}

type RestMiddlewareOpts struct {
	AuthConfig config.AuthorizationConfig
}

func NewRestMiddleware(o RestMiddlewareOpts) *RestMiddlewareModule {
	return &RestMiddlewareModule{
		authConfig: o.AuthConfig,
	}
}
