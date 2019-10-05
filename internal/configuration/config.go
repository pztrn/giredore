package configuration

import (
	// other
	"github.com/vrischmann/envconfig"
)

type config struct {
	HTTP struct {
		Listen         string `envconfig:"default=127.0.0.1:62222"`
		WaitForSeconds int    `envconfig:"default=10"`
	}
}

// Initialize loads configuration into memory.
func (cf *config) Initialize() {
	log.Info().Msg("Loading configuration...")

	_ = envconfig.Init(cf)

	log.Info().Msgf("Configuration parsed: %+v", cf)
}
