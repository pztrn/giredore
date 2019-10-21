package clientv1

import (
	// stdlib
	"strings"

	// local
	"go.dev.pztrn.name/giredore/internal/requester"
	"go.dev.pztrn.name/giredore/internal/structs"
)

func DeletePackage(args []string, options map[string]string) {
	req := &structs.PackageDeleteRequest{
		OriginalPath: args[0],
	}

	log.Info().Str("original path", req.OriginalPath).Msg("Sending package deletion request to giredored...")

	url := "http://" + options["server"] + "/_api/packages"

	data, err := requester.Delete(url, req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send package deletion request to giredored")
	}

	log.Debug().Msg("Got data: " + string(data))
}

func GetPackages(args []string, options map[string]string) {
	pkgs := strings.Split(args[0], ",")

	req := &structs.PackageGetRequest{}
	if pkgs[0] == "all" {
		req.All = true
	} else {
		req.PackageNames = pkgs
	}

	url := "http://" + options["server"] + "/_api/packages"

	log.Info().Msg("Getting packages data from giredore server...")

	data, err := requester.Post(url, req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get packages data from giredore server!")
	}

	log.Debug().Msg("Got data: " + string(data))
}

func SetPackage(args []string, options map[string]string) {
	pkg := &structs.Package{
		Description:  args[0],
		OriginalPath: args[1],
		RealPath:     args[2],
		VCS:          args[3],
	}

	// Execute some necessary checks.
	// If package's original path isn't starting with "/" - add it.
	if !strings.HasPrefix(pkg.OriginalPath, "/") {
		pkg.OriginalPath = "/" + pkg.OriginalPath
	}

	log.Info().Str("description", pkg.Description).Str("original path", pkg.OriginalPath).Str("real path", pkg.RealPath).Str("VCS", pkg.VCS).Msg("Sending set/update request to giredored...")

	url := "http://" + options["server"] + "/_api/packages"

	data, err := requester.Put(url, pkg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to send package update/set request to giredored")
	}

	log.Debug().Msg("Got data: " + string(data))
}
