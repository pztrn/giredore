package configuration

import (
	// local
	"go.dev.pztrn.name/giredore/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger

	envCfg *envConfig
	Cfg    *fileConfig
)

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "configuration").Logger()
	log.Info().Msg("Initializing...")

	envCfg = &envConfig{}
	envCfg.Initialize()

	Cfg = &fileConfig{}
	Cfg.Initialize()

	Cfg.HTTP.Listen = envCfg.HTTP.Listen
	Cfg.HTTP.WaitForSeconds = envCfg.HTTP.WaitForSeconds
}

func Shutdown() {
	Cfg.Save()
}
