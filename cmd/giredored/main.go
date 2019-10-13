package main

import (
	// stdlib
	"os"
	"os/signal"
	"syscall"

	// local
	serverv1 "go.dev.pztrn.name/giredore/domains/server/v1"
	"go.dev.pztrn.name/giredore/internal/configuration"
	"go.dev.pztrn.name/giredore/internal/httpserver"
	"go.dev.pztrn.name/giredore/internal/logger"
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
