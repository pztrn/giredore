package configuration

import (
	// other
	"github.com/vrischmann/envconfig"
)

// This structure represents configuration that will be parsed via
// environment variables. This configuration has higher priority
// than configuration loaded from file.
type envConfig struct {
	// DataDir is a directory where giredore will store it's data
	// like dynamic configuration file.
	DataDir string `envconfig:"default=/var/lib/giredore"`
	// HTTP describes HTTP server configuration.
	HTTP struct {
		// Listen is an address on which HTTP server will listen.
		Listen string `envconfig:"default=127.0.0.1:62222"`
		// WaitForSeconds is a timeout during which we will wait for
		// HTTP server be up. If timeout will pass and HTTP server won't
		// start processing requests - giredore will exit.
		WaitForSeconds int `envconfig:"default=10"`
	}
}

// Initialize parses environment variables into structure.
func (cf *envConfig) Initialize() {
	log.Info().Msg("Loading configuration...")

	_ = envconfig.Init(cf)

	log.Info().Msgf("Environment parsed: %+v", cf)
}
