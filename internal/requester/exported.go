package requester

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog"
	"go.dev.pztrn.name/giredore/internal/logger"
)

var log zerolog.Logger

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "requester").Logger()
	log.Info().Msg("Initializing...")
}

func Delete(url string, data interface{}) ([]byte, error) {
	return execRequest("DELETE", url, data)
}

// nolint:wrapcheck
func execRequest(method string, url string, data interface{}) ([]byte, error) {
	log.Debug().Str("method", method).Str("URL", url).Msg("Trying to execute HTTP request...")

	httpClient := getHTTPClient()

	var dataToSend []byte
	if data != nil {
		dataToSend, _ = json.Marshal(data)
	}

	// Compose HTTP request.
	httpReq, err := http.NewRequestWithContext(context.Background(), method, url, bytes.NewReader(dataToSend))
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

func Post(url string, data interface{}) ([]byte, error) {
	return execRequest("POST", url, data)
}

func Put(url string, data interface{}) ([]byte, error) {
	return execRequest("PUT", url, data)
}
