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
			ExpectContinueTimeout: time.Second * 5,
			DialContext: (&net.Dialer{
				Timeout: time.Second * 5,
			}).DialContext,
			ResponseHeaderTimeout: time.Second * 5,
			TLSHandshakeTimeout:   time.Second * 10,
		},
	}

	return c
}
