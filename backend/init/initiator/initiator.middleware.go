package initiator

import (
	"fudjie.waizly/backend-test/config"
	"fudjie.waizly/backend-test/init/service"
	"fudjie.waizly/backend-test/internal/middleware"
)

var (
	httpMiddleware middleware.RestMiddleware
)

func (i *Initiator) InitMiddleware(cfg *config.MainConfig) *service.Middleware {
	rmw := middleware.NewRestMiddleware(middleware.RestMiddlewareOpts{
		AuthConfig: cfg.Authorization,
	})

	return &service.Middleware{
		HttpMiddleware: rmw,
	}
}
