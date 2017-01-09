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

	pathList := []string{}
	for _, arg := range c.Args() {
		var path string

		if n, err := strconv.Atoi(arg); err == nil {
			pathList = PathList()
			if len(pathList) <= n {
				return fmt.Errorf("target index is not exists")
			}
			path = pathList[n] // The path is indicated by number
		} else {
			path = arg // Indicated by path name
		}

		os.RemoveAll(path)
		unregister(path)
	}

	return nil
}
