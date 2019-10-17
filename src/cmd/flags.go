package cmd

import (
	"github.com/serboox/os-cli/src/configs"
	"github.com/urfave/cli"
)

// GetFlags function define input cli params
func GetFlags(cliArgs *configs.CliArgs) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "auth-host",
			Usage:       "http://127.0.0.1:80/v3/auth/tokens",
			Destination: &cliArgs.AuthHost,
			EnvVar:      "OS_CLI_AUTH_HOST",
			//Value:       "http://10.0.2.15/identity/v3/auth/tokens",
		},
		cli.StringFlag{
			Name:        "login",
			Usage:       "User login",
			Destination: &cliArgs.Login,
			EnvVar:      "OS_CLI_USERLOGIN",
			//Value:       "admin",
		},
		cli.StringFlag{
			Name:        "pass",
			Usage:       "User password",
			Destination: &cliArgs.Password,
			EnvVar:      "OS_CLI_PASSWORD",
			//Value:       "pass1",
		},
		cli.StringFlag{
			Name:        "instName",
			Usage:       "Instance name",
			Destination: &cliArgs.InstName,
		},
		cli.StringFlag{
			Name:        "imageRef",
			Usage:       "Instance image",
			Destination: &cliArgs.ImageRef,
		},
		cli.Int64Flag{
			Name:        "flavorRef",
			Usage:       "Instance flavor",
			Destination: &cliArgs.FlavorRef,
		},
	}
}
