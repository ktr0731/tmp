package main

import "github.com/urfave/cli"

var commands = []cli.Command{
	{
		Name:      "make",
		Aliases:   []string{"m", "mk"},
		Usage:     "Make a new temporary directory",
		ArgsUsage: "[-n] [name]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name, n",
				Value: "tmp",
				Usage: "Directory `name`",
			},
		},
		Action: makeDir,
	},
	{
		Name:    "remove",
		Aliases: []string{"r", "rm"},
		Usage:   "Remove a target directory as passed by an argument",
	},
	{
		Name:    "list",
		Aliases: []string{"l", "ls"},
		Usage:   "Show all tracking temporary directories",
		Action:  list,
	},
}
