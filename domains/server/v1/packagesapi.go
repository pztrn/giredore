package serverv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/structs"
)

// This function responsible for getting packages configuration.
func packagesGET(ectx echo.Context) error {
	// nolint:exhaustruct
	req := &structs.PackageGetRequest{}
	if err := ectx.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package get request")

		// nolint:exhaustruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingPackagesGetRequest}})
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
		// nolint:wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errors, Data: pkgs})
	}

	// nolint:exhaustruct,wrapcheck
	return ectx.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess, Data: pkgs})
}

// This function responsible for deleting package.
func packagesDELETE(ectx echo.Context) error {
	// nolint:exhaustruct
	req := &structs.PackageDeleteRequest{}
	if err := ectx.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package delete request")

		// nolint:exhaustruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrParsingDeleteRequest}})
	}

	log.Info().Msgf("Received package delete request: %+v", req)

	errs := configuration.Cfg.DeletePackage(req)

	if len(errs) > 0 {
		// nolint:exhaustruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errs})
	}

	// nolint:exhaustruct,wrapcheck
	return ectx.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}

// This function responsible for setting or updating packages.
func packagesSET(ectx echo.Context) error {
	// nolint:exhaustruct
	req := &structs.Package{}
	if err := ectx.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to parse package data")

		// nolint:wrapcheck
		return ectx.JSON(http.StatusBadRequest, nil)
	}

	log.Info().Msgf("Received package set/update request: %+v", req)

	// Validate passed package data.
	if !strings.HasPrefix(req.OriginalPath, "/") {
		// nolint:exhaustruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrPackageOrigPathShouldStartWithSlash}})
	}

	configuration.Cfg.AddOrUpdatePackage(req)

	// nolint:exhaustruct,wrapcheck
	return ectx.JSON(http.StatusOK, &structs.Reply{Status: structs.StatusSuccess})
}
