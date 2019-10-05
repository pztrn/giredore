package clientv1

import (
	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/requester"
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
