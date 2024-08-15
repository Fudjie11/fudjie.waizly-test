package server

import (
	"fudjie.waizly/backend-test/init/service"

	_ "fudjie.waizly/backend-test/docs" // swagger docs
	"fudjie.waizly/backend-test/library/env"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	employeeRouter "fudjie.waizly/backend-test/internal/app/main/v1/employee/rest/router"
)

type Options struct {
	DefaultTimeout int
	Service        *service.RESTService
}

type Server struct {
	Options *Options
	Service *service.RESTService
}

func NewRESTServer(o *Options) *Server {
	return &Server{
		Options: o,
		Service: o.Service,
	}
}

func (s *Server) Serve(srv *fiber.App) {
	srv.Use(logger.New())
	srv.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "*",
	}))

	s.initSwagger(srv)
	s.initRoutes(srv)
}

func (s *Server) initSwagger(srv *fiber.App) {
	if s.Service.Config.Server.Rest.EnableSwagger {
		if env.GetEnvironmentName() != "production" {
			srv.Use("swagger", swagger.HandlerDefault)
		}
	}
}

func (s *Server) initRoutes(srv *fiber.App) {
	apiMainV1 := srv.Group("/api/main/v1")
	httpMiddleware := s.Service.Middleware.HttpMiddleware
	restHandler := s.Service.RESTHandler

	employeeRouter.Configure(apiMainV1, httpMiddleware, restHandler.EmployeeRestHandler)
}
