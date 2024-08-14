package servers

import (
	"fmt"
	"time"

	api "fudjie.waizly/backend-test/handler/rest/server"
	"fudjie.waizly/backend-test/init/service"

	"github.com/gofiber/fiber/v2"
)

type Opts struct {
	ListenAddress string
	Port          int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	Service       *service.RESTService
}

type Handler struct {
	options     *Opts
	server      *fiber.App
	listenErrCh chan error
}

type RESTHandler interface {
	RunREST()
	ListenError() <-chan error
}

func NewREST(o *Opts) RESTHandler {
	srv := fiber.New(fiber.Config{
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	})

	// Liveness check api
	srv.Get("/health", func(fc *fiber.Ctx) error {
		return fc.JSON("OK")
	})

	api.NewRESTServer(&api.Options{
		DefaultTimeout: o.Service.Config.Server.Rest.DefaultTimeout,
		Service:        o.Service,
	}).Serve(srv)

	handler := &Handler{
		options:     o,
		server:      srv,
		listenErrCh: make(chan error),
	}

	return handler
}

func (h *Handler) RunREST() {
	listenAddress := fmt.Sprintf("%s:%d", h.options.ListenAddress, h.options.Port)
	h.listenErrCh <- h.server.Listen(listenAddress)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
