package badge

import (
	"github.com/codegangsta/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Adds a new Host",
			Action: func(c *cli.Context) {
				println("Host added: ", c.Args().First())
			},
		},
	}
}
