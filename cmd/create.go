package cmd

import (
	"github.com/serboox/os-cli/configs"
	"github.com/serboox/os-cli/controllers"
	"github.com/urfave/cli"
)

//CreateInstance cli command
func CreateInstance(
	flags []cli.Flag,
	cliArgs *configs.CliArgs,
) cli.Command {
	return cli.Command{
		Name:    "create",
		Aliases: []string{"create"},
		Usage:   "Create instance",
		Flags:   flags,
		Action: func(c *cli.Context) error {
			return controllers.CreateInstance(
				cliArgs,
			)
		},
	}
}
