package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func makeDir(c *cli.Context) error {
	dir := "."
	if c.NArg() == 1 {
		dir = c.Args()[0]
	} else if c.NArg() != 0 {
		return fmt.Errorf("invalid arguments\nUsage: tmp make " + c.Command.ArgsUsage)
	}

	name := c.String("name")

	fmt.Println("directory: " + dir + " name: " + name)

	return nil
}
