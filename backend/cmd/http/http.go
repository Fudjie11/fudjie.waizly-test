package http

import (
	"fmt"
	"log"

	"fudjie.waizly/backend-test/init/assembler"

	"fudjie.waizly/backend-test/library/env"
	"fudjie.waizly/backend-test/library/logger"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:              "serve-http",
		Short:            "Serve REST API",
		Long:             "Serve REST API",
		PersistentPreRun: rootRESTPreRun,
		RunE:             runREST,
	}
	serviceName = fmt.Sprintf("%s-bluebird.tech/tms-rest", env.GetEnvironmentName())
)

func rootRESTPreRun(cmd *cobra.Command, args []string) {
	// initiate logger
	logger.InitGlobalLogger(&logger.Config{
		ServiceName: "bluebird-tms",
		Level:       zerolog.DebugLevel,
	})
}

func ServeHTTPCmd() *cobra.Command {
	return serveHTTPCmd
}

func runREST(cmd *cobra.Command, args []string) error {
	configURL, _ := cmd.Flags().GetString("config")
	bootstrapREST(assembler.NewAssembler(), configURL)
	return nil
}

func bootstrapREST(starter assembler.AssemblerManager, configPath string) {

	starter = starter.BuildService(configPath, serviceName).AssembleRESTApplication()
	starter.RunRESTApplication()

	select {
	case err := <-starter.ListenErrorRESTApp():
		log.Fatalf("Error starting rest server, exiting gracefully %v:", err)
	case <-starter.TerminateSignal():
		log.Fatalln("Exiting gracefully...")
	}
}
