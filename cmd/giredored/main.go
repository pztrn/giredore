package main

import (
	// stdlib
	"os"
	"os/signal"
	"syscall"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/domains/server/v1"
	"sources.dev.pztrn.name/pztrn/giredore/internal/configuration"
	"sources.dev.pztrn.name/pztrn/giredore/internal/httpserver"
	"sources.dev.pztrn.name/pztrn/giredore/internal/logger"
)

func main() {
	logger.Initialize()
	logger.Logger.Info().Msg("Starting giredore server...")

	configuration.Initialize()
	httpserver.Initialize()

	serverv1.Initialize()

	httpserver.Start()

	logger.Logger.Info().Msg("giredore server is ready")

	// CTRL+C handler.
	signalHandler := make(chan os.Signal, 1)
	shutdownDone := make(chan bool, 1)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalHandler
		httpserver.Shutdown()
		configuration.Shutdown()
		shutdownDone <- true
	}()

	<-shutdownDone
	os.Exit(0)
}
