package serverv1

import (
	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/httpserver"
	"sources.dev.pztrn.name/pztrn/giredore/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
)

func Initialize() {
	log = logger.Logger.With().Str("type", "domain").Str("package", "server").Int("version", 1).Logger()
	log.Info().Msg("Initializing...")

	httpserver.Srv.GET("/_api/configuration", configurationGET)
}
