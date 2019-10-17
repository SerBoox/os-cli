package cmd

import (
	"github.com/serboox/os-cli/src/configs"
	"github.com/serboox/os-cli/src/controllers"
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
		Action: func(c *cli.Context) (err error) {
			err = cliArgs.ValidateAuthHost()
			if err != nil {
				return err
			}
			err = cliArgs.ValidateLogin()
			if err != nil {
				return err
			}
			err = cliArgs.ValidatePassword()
			if err != nil {
				return err
			}

			return controllers.ShowInstances(
				cliArgs,
			)
		},
	}
}
