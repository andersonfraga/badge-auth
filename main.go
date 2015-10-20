package main

import (
	"github.com/andersonfraga/badge-auth/badge"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Badge Auth"
	app.Usage = ""
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Commands = badge.GetCommands()

	app.Run(os.Args)
}
