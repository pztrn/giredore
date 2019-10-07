package main

import (
	// stdlib
	"os"

	// local
	clientv1 "sources.dev.pztrn.name/pztrn/giredore/domains/client/v1"
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
		).
		WithCommand(
			cli.NewCommand("set", "sets configuration value").
				WithCommand(
					cli.NewCommand("allowedips", "sets list of allowed IPs for interacting with configuration API").
						WithArg(cli.NewArg("allowed_ips_list", "list of allowed IP addresses delimited by comma. Subnets are also fine.")).
						WithAction(func(args []string, options map[string]string) int {
							logger.Initialize()
							clientv1.Initialize()
							clientv1.SetAllowedIPs(args, options)
							return 0
						}),
				),
		)

	packages := cli.NewCommand("packages", "work with packages giredore will serve").
		WithShortcut("pkg").
		WithCommand(
			cli.NewCommand("get", "gets and prints out list of packages that is served by giredore").
				WithArg(cli.NewArg("pkgnames", "one or more packages to get info about, delimited with comma, e.g. '/path/pkg1,/path/pkg2'. Say 'all' here to get info about all known packages.")).
				WithAction(func(args []string, options map[string]string) int {
					logger.Initialize()
					clientv1.Initialize()
					clientv1.GetPackages(args, options)

					return 0
				}),
		).
		WithCommand(
			cli.NewCommand("set", "creates or updates package data").
				WithArg(cli.NewArg("description", "optional package description that will be shown on package serving page")).
				WithArg(cli.NewArg("origpath", "original path of package without domain, e.g. '/group/pkg' instead of 'github.com/group/pkg'")).
				WithArg(cli.NewArg("realpath", "real path for package sources, e.g. 'github.com/group/pkg.git'")).
				WithArg(cli.NewArg("vcs", "VCS used for package sources getting. See https://github.com/golang/tools/blob/master/go/vcs/vcs.go for list of supported VCS.")).
				WithAction(func(args []string, options map[string]string) int {
					logger.Initialize()
					clientv1.Initialize()
					clientv1.SetPackage(args, options)

					return 0
				}),
		).
		WithCommand(
			cli.NewCommand("delete", "deletes package data").
				WithArg(cli.NewArg("origpath", "original path of package without domain, e.g. '/group/pkg' instead of 'github.com/group/pkg'")).
				WithAction(func(args []string, options map[string]string) int {
					logger.Initialize()
					clientv1.Initialize()
					clientv1.DeletePackage(args, options)

					return 0
				}),
		)

	app := cli.New("giredore server controlling utility").
		WithOption(
			cli.NewOption("server", "giredore server address"),
		).
		WithCommand(config).
		WithCommand(packages)

	os.Exit(app.Run(os.Args, os.Stdout))
}
