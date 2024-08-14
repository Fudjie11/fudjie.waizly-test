package initiator

import (
	"fudjie.waizly/backend-test/config"

	"github.com/google/gops/agent"
	"github.com/rs/zerolog/log"
)

var (
	errInitConfig = "failed to initiate config"
)

// Main Config
func (i *Initiator) InitConfig(configPath string, serviceName string) *config.MainConfig {
	if err := i.AgentListen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	cfg := &config.MainConfig{}
	log.Info().Msgf("reading config from %s", configPath)
	err := i.ReadConfig(cfg, configPath, "config")
	if err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	return cfg

}
