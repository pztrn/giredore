package serverv1

import (
	// stdlib
	"net/http"
	"strings"

	// local
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/structs"

	// other
	"github.com/labstack/echo"
)

// This function responsible for getting packages configuration.
func packagesGET(ec echo.Context) error {
	req := &structs.PackageGetRequest{}
	if err := ec.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package get request")
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingPackagesGetRequest}})
	}

	log.Info().Msgf("Received package(s) info get request: %+v", req)

	var pkgs map[string]*structs.Package

	var errors []structs.Error

	if req.All {
		pkgs = configuration.Cfg.GetAllPackagesInfo()
	} else {
		pkgs, errors = configuration.Cfg.GetPackagesInfo(req.PackageNames)
	}

	if len(errors) > 0 {
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errors, Data: pkgs})
	}

	return ec.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess, Data: pkgs})
}

// This function responsible for deleting package.
func packagesDELETE(ec echo.Context) error {
	req := &structs.PackageDeleteRequest{}
	if err := ec.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package delete request")
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingDeleteRequest}})
	}

	log.Info().Msgf("Received package delete request: %+v", req)

	errs := configuration.Cfg.DeletePackage(req)

	if len(errs) > 0 {
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errs})
	}

	return ec.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}

// This function responsible for setting or updating packages.
func packagesSET(ec echo.Context) error {
	req := &structs.Package{}
	if err := ec.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package data")
		return ec.JSON(http.StatusBadRequest, nil)
	}

	log.Info().Msgf("Received package set/update request: %+v", req)

	// Validate passed package data.
	if !strings.HasPrefix(req.OriginalPath, "/") {
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrPackageOrigPathShouldStartWithSlash}})
	}

	configuration.Cfg.AddOrUpdatePackage(req)

	return ec.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}
