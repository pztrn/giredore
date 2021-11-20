package clientv1

import (
	"strings"

	"go.dev.pztrn.name/giredore/internal/requester"
	"go.dev.pztrn.name/giredore/internal/structs"
)

func GetConfiguration(options map[string]string) {
	url := "http://" + options["server"] + "/_api/configuration"

	log.Info().Msg("Getting configuration from giredore server...")

	data, err := requester.Get(url)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get configuration from giredore server!")
	}

	log.Debug().Msg("Got data: " + string(data))
}

func SetAllowedIPs(args []string, options map[string]string) {
	url := "http://" + options["server"] + "/_api/configuration/allowedips"

	log.Info().Str("allowed IPs", args[0]).Msg("Setting allowed IPs for API interaction...")

	req := &structs.AllowedIPsSetRequest{
		AllowedIPs: strings.Split(args[0], ","),
	}

	data, err := requester.Post(url, req)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set allowed IPs in giredore server configuration!")
	}

	log.Debug().Msg("Got data: " + string(data))
}
