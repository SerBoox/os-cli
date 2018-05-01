package cmd

import (
	"github.com/serboox/os-cli/configs"
	"github.com/serboox/os-cli/controllers"
	"github.com/urfave/cli"
)

// ShowInstances cli command
func ShowInstances(
	flags []cli.Flag,
	cliArgs *configs.CliArgs,
) cli.Command {
	return cli.Command{
		Name:    "show",
		Aliases: []string{"show"},
		Usage:   "Show instances",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			return controllers.ShowInstances(
				cliArgs,
			)
		},
	}
}
