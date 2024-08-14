package assembler

import (
	"time"

	mainREST "fudjie.waizly/backend-test/handler/rest"
	"fudjie.waizly/backend-test/init/service"
)

func (a *assembler) assembleREST(service *service.RESTService) mainREST.RESTHandler {
	server := a.NewRESTHandler(&mainREST.Opts{
		Service:       service,
		ListenAddress: service.Config.Server.Rest.ListenAddress,
		Port:          service.Config.Server.Rest.Port,
		ReadTimeout:   time.Duration(service.Config.Server.Rest.ReadTimeout) * time.Millisecond,
		WriteTimeout:  time.Duration(service.Config.Server.Rest.WriteTimeout) * time.Millisecond,
	})

	return server
}

func (a *assembler) runRESTServer() {
	go a.restHandler.RunREST()
}
