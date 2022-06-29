package serverv1

import (
	"net/http"

	"github.com/labstack/echo"
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/structs"
)

// This function responsible for getting runtime configuration.
func configurationGET(ec echo.Context) error {
	// nolint:wrapcheck
	return ec.JSON(http.StatusOK, configuration.Cfg)
}

func configurationAllowedIPsSET(ectx echo.Context) error {
	// nolint:exhaustruct
	req := &structs.AllowedIPsSetRequest{}
	if err := ectx.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse allowed IPs set request")
		// nolint:exhaustruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingAllowedIPsSetRequest}})
	}

	log.Debug().Msgf("Got set allowed IPs request: %+v", req)

	configuration.Cfg.SetAllowedIPs(req.AllowedIPs)

	// nolint:exhaustruct,wrapcheck
	return ectx.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}
