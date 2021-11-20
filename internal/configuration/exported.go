package configuration

import (
	"github.com/rs/zerolog"
	"go.dev.pztrn.name/giredore/internal/logger"
)

var (
	log zerolog.Logger

	envCfg *envConfig
	Cfg    *fileConfig
)

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "configuration").Logger()
	log.Info().Msg("Initializing...")

	// nolint:exhaustivestruct
	envCfg = &envConfig{}
	envCfg.Initialize()

	// nolint:exhaustivestruct
	Cfg = &fileConfig{}
	Cfg.Initialize()

	Cfg.HTTP.Listen = envCfg.HTTP.Listen
	Cfg.HTTP.WaitForSeconds = envCfg.HTTP.WaitForSeconds
}

func Shutdown() {
	Cfg.Save()
}
