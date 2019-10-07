package serverv1

import (
	// stdlib
	"net/http"
	"strings"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/configuration"
	"sources.dev.pztrn.name/pztrn/giredore/internal/structs"

	// other
	"github.com/labstack/echo"
)

func throwGoImports(ec echo.Context) error {
	// Getting real path. This might be the package itself, or namespace
	// to list available packages.
	// For now only package itself is supported, all other features in ToDo.
	packageNameRaw := ec.Request().URL.Path
	pkgs, errs := configuration.Cfg.GetPackagesInfo([]string{packageNameRaw})
	if errs != nil {
		log.Error().Str("package", packageNameRaw).Msgf("Failed to get package information: %+v", errs)
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: errs})
	}

	if len(pkgs) == 0 {
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrNoPackagesFound}})
	}

	pkg, found := pkgs[packageNameRaw]
	if !found {
		return ec.JSON(http.StatusBadRequest, &structs.Reply{Status: structs.StatusFailure, Errors: []structs.Error{structs.ErrNoPackagesFound}})
	}

	// We should compose package name using our domain under which giredore
	// is working.
	domain := ec.Request().Host
	packageName := domain + packageNameRaw

	tmpl := singlePackageTemplate
	tmpl = strings.Replace(tmpl, "{PKGNAME}", packageName, -1)
	tmpl = strings.Replace(tmpl, "{VCS}", pkg.VCS, 1)
	tmpl = strings.Replace(tmpl, "{REPOPATH}", pkg.RealPath, 1)

	return ec.HTML(http.StatusOK, tmpl)
}
