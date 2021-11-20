package serverv1

import (
	"github.com/rs/zerolog"
	"go.dev.pztrn.name/giredore/internal/httpserver"
	"go.dev.pztrn.name/giredore/internal/logger"
)

var log zerolog.Logger

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
