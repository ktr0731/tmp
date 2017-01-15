package main

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli"
)

func move(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("invalid arguments")
	}

	path, err := filepath.Abs(c.Args()[0])
	if err != nil {
		return fmt.Errorf("cannot convert path to absolute: %s", err)
	}

	if contains(path) {
		unregister(c.Args()[0])
		register(c.Args()[1])
	} else {
		return fmt.Errorf("passed path is not exists: %s", c.Args()[0])
	}

	return nil
}
