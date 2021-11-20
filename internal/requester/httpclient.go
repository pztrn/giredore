package requester

import (
	"net"
	"net/http"
	"time"
)

// nolint:exhaustivestruct
func getHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			ExpectContinueTimeout: time.Second * 5,
			DialContext: (&net.Dialer{
				Timeout: time.Second * 5,
			}).DialContext,
			ResponseHeaderTimeout: time.Second * 5,
			TLSHandshakeTimeout:   time.Second * 10,
		},
	}

	return client
}
