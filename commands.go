package main

import "github.com/urfave/cli"

var commands = []cli.Command{
	{
		Name:    "make",
		Aliases: []string{"m", "mk"},
		Usage:   "Make a new temporary directory",
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
	},
}
