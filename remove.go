package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func removeDir(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("invalid argument")
	}

	for _, path := range c.Args() {
		if n, err := strconv.Atoi(path); err == nil {
			os.RemoveAll(PathList()[n]) // The path is indicated by number
		} else {
			os.RemoveAll(path) // Indicated by path name
		}
	}

	return nil
}
