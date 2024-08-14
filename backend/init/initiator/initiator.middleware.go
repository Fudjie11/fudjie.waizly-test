package initiator

import (
	"fudjie.waizly/backend-test/init/service"
	"fudjie.waizly/backend-test/internal/middleware"
)

var (
	httpMiddleware middleware.HttpMiddleware
)

func (i *Initiator) InitMiddleware() *service.Middleware {
	httpMiddleware = middleware.NewHttpMiddleware()

	return &service.Middleware{
		HttpMiddleware: httpMiddleware,
	}
}
