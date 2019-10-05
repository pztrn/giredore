package main

import (
	// stdlib
	"os"

	// local
	"sources.dev.pztrn.name/pztrn/giredore/domains/client/v1"
	"sources.dev.pztrn.name/pztrn/giredore/internal/logger"

	// other
	"github.com/teris-io/cli"
)

func main() {
	config := cli.NewCommand("configuration", "work with giredore server configuration").
		WithShortcut("conf").
		WithCommand(
			cli.NewCommand("get", "gets and prints out current giredore server configuration").
				WithAction(func(args []string, options map[string]string) int {
					logger.Initialize()
					clientv1.Initialize()
					clientv1.GetConfiguration(options)

					return 0
				}),
		)

	app := cli.New("giredore server controlling utility").
		WithOption(
			cli.NewOption("server", "giredore server address"),
		).
		WithCommand(config)

	os.Exit(app.Run(os.Args, os.Stdout))
}
