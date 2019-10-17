package cmd

import (
	"github.com/serboox/os-cli/src/configs"
	"github.com/serboox/os-cli/src/controllers"
	"github.com/urfave/cli"
)

// CreateInstance cli command
func CreateInstance(
	flags []cli.Flag,
	cliArgs *configs.CliArgs,
) cli.Command {
	return cli.Command{
		Name:    "create",
		Aliases: []string{"create"},
		Usage:   "Create instance",
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
			err = cliArgs.ValidateInstName()
			if err != nil {
				return err
			}
			err = cliArgs.ValidateImageRef()
			if err != nil {
				return err
			}
			err = cliArgs.ValidateFlavorRef()
			if err != nil {
				return err
			}

			return controllers.CreateInstance(
				cliArgs,
			)
		},
	}
}
