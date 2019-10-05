package requester

import (
	// stdlib
	"net"
	"net/http"
	"time"
)

func getHTTPClient() *http.Client {
	c := &http.Client{
		Transport: &http.Transport{
			// ToDo: configurable.
			ExpectContinueTimeout: time.Second * 5,
			DialContext: (&net.Dialer{
				// ToDo: configurable.
				Timeout: time.Second * 5,
			}).DialContext,
			// ToDo: configurable.
			ResponseHeaderTimeout: time.Second * 5,
			// ToDo: configurable.
			TLSHandshakeTimeout: time.Second * 10,
		},
	}

	// ToDo: skip verifying insecure certificates if option was
	// specified.

	return c
}
