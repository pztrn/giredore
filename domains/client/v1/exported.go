package clientv1

import (
	"github.com/rs/zerolog"
	"go.dev.pztrn.name/giredore/internal/logger"
	"go.dev.pztrn.name/giredore/internal/requester"
)

var log zerolog.Logger

func Initialize() {
	log = logger.Logger.With().Str("type", "domain").Str("package", "client").Int("version", 1).Logger()
	log.Info().Msg("Initializing...")

	requester.Initialize()
}
