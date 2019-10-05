package httpserver

import (
	// stdlib
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/internal/configuration"
	"sources.dev.pztrn.name/pztrn/giredore/internal/logger"

	// other
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger

	Srv *echo.Echo
)

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "httpserver").Logger()

	log.Info().Msg("Initializing...")

	Srv = echo.New()
	Srv.Use(middleware.Recover())
	Srv.Use(requestLogger())
	Srv.DisableHTTP2 = true
	Srv.HideBanner = true
	Srv.HidePort = true

	Srv.GET("/_internal/waitForOnline", waitForHTTPServerToBeUpHandler)
}

// Shutdown stops HTTP server. Returns true on success and false on failure.
func Shutdown() {
	log.Info().Msg("Shutting down HTTP server...")
	err := Srv.Shutdown(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to stop HTTP server")
	}
	log.Info().Msg("HTTP server shutted down")

}

// Start starts HTTP server and checks that server is ready to process
// requests. Returns true on success and false on failure.
func Start() {
	log.Info().Str("address", configuration.Cfg.HTTP.Listen).Msg("Starting HTTP server...")

	go func() {
		err := Srv.Start(configuration.Cfg.HTTP.Listen)
		if !strings.Contains(err.Error(), "Server closed") {
			log.Fatal().Err(err).Msg("HTTP server critial error occured")
		}
	}()

	// Check that HTTP server was started.
	httpc := &http.Client{Timeout: time.Second * 1}
	checks := 0
	for {
		checks++
		if checks >= configuration.Cfg.HTTP.WaitForSeconds {
			log.Fatal().Int("seconds passed", checks).Msg("HTTP server isn't up")
		}
		time.Sleep(time.Second * 1)
		resp, err := httpc.Get("http://" + configuration.Cfg.HTTP.Listen + "/_internal/waitForOnline")
		if err != nil {
			log.Debug().Err(err).Msg("HTTP error occured, HTTP server isn't ready, waiting...")
			continue
		}

		response, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Debug().Err(err).Msg("Failed to read response body, HTTP server isn't ready, waiting...")
			continue
		}
		log.Debug().Str("status", resp.Status).Int("body length", len(response)).Msg("HTTP response received")

		if resp.StatusCode == http.StatusOK {
			if len(response) == 0 {
				log.Debug().Msg("Response is empty, HTTP server isn't ready, waiting...")
				continue
			}
			log.Debug().Int("status code", resp.StatusCode).Msgf("Response: %+v", string(response))
			if len(response) == 17 {
				break
			}
		}
	}
	log.Info().Msg("HTTP server is ready to process requests")
}

func waitForHTTPServerToBeUpHandler(ec echo.Context) error {
	response := map[string]string{
		"error": "None",
	}
	return ec.JSON(200, response)
}
