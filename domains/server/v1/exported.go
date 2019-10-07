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

	// Configuration-related.
	httpserver.Srv.GET("/_api/configuration", configurationGET)
	httpserver.Srv.POST("/_api/configuration/allowedips", configurationAllowedIPsSET)

	// Packages-related.
	httpserver.Srv.POST("/_api/packages", packagesGET)
	httpserver.Srv.PUT("/_api/packages", packagesSET)
	httpserver.Srv.DELETE("/_api/packages", packagesDELETE)

	// goimports serving.
	httpserver.Srv.GET("/*", throwGoImports)
}
