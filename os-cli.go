package main

import (
	"fmt"
	"os"

	"github.com/serboox/os-cli/cmd"
	"github.com/serboox/os-cli/configs"
	"github.com/urfave/cli"
)

var (
	cliArgs configs.CliArgs
)

func main() {
	app := cli.NewApp()
	app.Name = "os-cli"
	app.Usage = "OpenStack Client"
	app.Version = "1.0.0"

	flags := cmd.GetFlags(&cliArgs)

	app.Commands = []cli.Command{
		cmd.CreateInstance(flags, &cliArgs),
		cmd.ShowInstances(flags, &cliArgs),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
