package httpserver

import (
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/structs"
)

func checkAllowedIPs() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			// Do nothing if request came not in "/_api" namespace.
			if !strings.HasPrefix(ectx.Request().RequestURI, "/_api") {
				return next(ectx)
			}

			// Get IPs and subnets from configuration and parse them
			// into comparable things.
			// If IP address was specified without network mask - assume /32.
			var subnets []*net.IPNet

			allowedIPs := configuration.Cfg.GetAllowedIPs()

			for _, ip := range allowedIPs {
				ipToParse := ip
				if !strings.Contains(ip, "/") {
					ipToParse = ip + "/32"
				}

				_, net, err := net.ParseCIDR(ipToParse)
				if err != nil {
					log.Error().Err(err).Str("subnet", ipToParse).Msg("Failed to parse CIDR. /_api/ endpoint won't be accessible, this should be fixed manually in configuration file!")

					// nolint:exhaustivestruct,wrapcheck
					return ectx.JSON(http.StatusInternalServerError, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrInvalidAllowedIPDefined}})
				}

				subnets = append(subnets, net)
			}

			// Check if requester's IP address are within allowed IP
			// subnets.
			ipToCheck := net.ParseIP(ectx.RealIP())

			var allowed bool

			for _, subnet := range subnets {
				if subnet.Contains(ipToCheck) {
					allowed = true

					break
				}
			}

			if allowed {
				return next(ectx)
			}

			// nolint:exhaustivestruct,wrapcheck
			return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrIPAddressNotAllowed}})
		}
	}
}
