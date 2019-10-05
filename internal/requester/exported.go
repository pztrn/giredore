package requester

import (
	// stdlib
	"io/ioutil"
	"net/http"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
)

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "requester").Logger()
	log.Info().Msg("Initializing...")
}

func execRequest(method string, url string, data map[string]string) ([]byte, error) {
	log.Debug().Str("method", method).Str("URL", url).Msg("Trying to execute HTTP request...")

	httpClient := getHTTPClient()

	// Compose HTTP request.
	// ToDo: POST/PUT/other methods that require body.
	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	response, err1 := httpClient.Do(httpReq)
	if err1 != nil {
		return nil, err1
	}

	bodyBytes, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return nil, err2
	}
	response.Body.Close()

	log.Debug().Int("response body length (bytes)", len(bodyBytes)).Msg("Got response")

	return bodyBytes, nil
}

func Get(url string) ([]byte, error) {
	return execRequest("GET", url, nil)
}
