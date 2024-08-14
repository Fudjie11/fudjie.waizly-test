package assembler

import (
	"os"
	"os/signal"
	"syscall"

	handler_rest "fudjie.waizly/backend-test/handler/rest"

	"fudjie.waizly/backend-test/init/initiator"
	"fudjie.waizly/backend-test/init/service"
)

type NewRESTHandler func(o *handler_rest.Opts) handler_rest.RESTHandler

type assembler struct {
	Initiator initiator.InitiatorManager
	term      chan os.Signal

	NewRESTHandler NewRESTHandler
	restHandler    handler_rest.RESTHandler
	restService    *service.RESTService
}

type AssemblerManager interface {
	BuildService(configPath string, ServiceName string) AssemblerManager

	AssembleRESTApplication() AssemblerManager
	RunRESTApplication()
	ListenErrorRESTApp() <-chan error

	TerminateSignal() chan os.Signal
}

func NewAssembler() AssemblerManager {
	return &assembler{
		Initiator:      initiator.New(),
		NewRESTHandler: handler_rest.NewREST,
	}
}

func (a *assembler) BuildService(configPath string, serviceName string) AssemblerManager {
	cfg := a.Initiator.InitConfig(configPath, serviceName)
	infra := a.Initiator.InitInfrastructure(cfg)
	repo := a.Initiator.InitRepository(cfg, infra)
	uc := a.Initiator.InitUsecase(cfg, infra, repo)
	middleware := a.Initiator.InitMiddleware()

	hdlr := a.Initiator.InitRESTHandler(cfg, infra, uc)
	restsvc := a.Initiator.InitRESTService(cfg, infra, middleware, hdlr, uc)
	a.restService = restsvc

	return a
}

func (a *assembler) AssembleRESTApplication() AssemblerManager {
	a.restHandler = a.assembleREST(a.restService)
	return a
}

func (a *assembler) RunRESTApplication() {
	a.runRESTServer()
}

func (a *assembler) ListenErrorRESTApp() <-chan error {
	return a.restHandler.ListenError()
}

func (a *assembler) TerminateSignal() chan os.Signal {
	a.term = make(chan os.Signal)
	signal.Notify(a.term, os.Interrupt, syscall.SIGTERM)
	return a.term
}
