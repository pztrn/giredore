package serverv1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/structs"
)

func throwGoImports(ectx echo.Context) error {
	// Getting real path. This might be the package itself, or namespace
	// to list available packages.
	// For now only package itself is supported, all other features in ToDo.
	packageNameRaw := ectx.Request().URL.Path

	pkgs, errs := configuration.Cfg.GetPackagesInfo([]string{packageNameRaw})

	if errs != nil {
		log.Error().Str("package", packageNameRaw).Msgf("Failed to get package information: %+v", errs)

		// nolint:exhaustivestruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errs})
	}

	if len(pkgs) == 0 {
		// nolint:exhaustivestruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrNoPackagesFound}})
	}

	pkg, found := pkgs[packageNameRaw]
	if !found {
		// nolint:exhaustivestruct,wrapcheck
		return ectx.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrNoPackagesFound}})
	}

	// We should compose package name using our domain under which giredore
	// is working.
	domain := ectx.Request().Host
	packageName := domain + packageNameRaw

	tmpl := singlePackageTemplate
	tmpl = strings.Replace(tmpl, "{PKGNAME}", packageName, -1)
	tmpl = strings.Replace(tmpl, "{VCS}", pkg.VCS, 1)
	tmpl = strings.Replace(tmpl, "{REPOPATH}", pkg.RealPath, 1)

	// nolint:wrapcheck
	return ectx.HTML(http.StatusOK, tmpl)
}
