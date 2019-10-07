package httpserver

import (
	// stdlib
	"net"
	"net/http"
	"strings"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/configuration"
	"sources.dev.pztrn.name/pztrn/giredore/internal/structs"

	// other
	"github.com/labstack/echo"
)

func checkAllowedIPs() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			// Do nothing if request came not in "/_api" namespace.
			if !strings.HasPrefix(ec.Request().RequestURI, "/_api") {
				_ = next(ec)
				return nil
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
					return ec.JSON(http.StatusInternalServerError, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrInvalidAllowedIPDefined}})
				}

				subnets = append(subnets, net)
			}

			// Check if requester's IP address are within allowed IP
			// subnets.
			ipToCheck := net.ParseIP(ec.RealIP())
			var allowed bool
			for _, subnet := range subnets {
				if subnet.Contains(ipToCheck) {
					allowed = true
					break
				}
			}

			if allowed {
				_ = next(ec)
				return nil
			}

			return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrIPAddressNotAllowed}})
		}
	}
}
