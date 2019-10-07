package serverv1

import (
	// stdlib
	"net/http"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/configuration"
	"sources.dev.pztrn.name/pztrn/giredore/internal/structs"

	// other
	"github.com/labstack/echo"
)

// This function responsible for getting runtime configuration.
func configurationGET(ec echo.Context) error {
	return ec.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func configurationAllowedIPsSET(ec echo.Context) error {
	req := &structs.AllowedIPsSetRequest{}
	if err := ec.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse allowed IPs set request")
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingAllowedIPsSetRequest}})
	}

	log.Debug().Msgf("Got set allowed IPs request: %+v", req)

	configuration.Cfg.SetAllowedIPs(req.AllowedIPs)

	return ec.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}
