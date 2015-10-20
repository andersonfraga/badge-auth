package main 

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/btfidelis/badge-auth/badge"
)

func main() {
	app := cli.NewApp()
	app.Name = "SKeyL"
	app.Usage = ""
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	app.Commands = badge.GetCommands()

	app.Run(os.Args)
}
