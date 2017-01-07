package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func removeDir(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("invalid argument")
	}

	for _, path := range c.Args() {
		os.RemoveAll(path)
	}

	return nil
}
