package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func cleanDir(path string) {
	os.RemoveAll(path)
	unregister(path)
}

func removeDir(c *cli.Context) error {
	pathList := []string{}

	if c.NArg() == 0 {
		if c.Bool("all") {
			for _, path := range PathList() {
				cleanDir(path)
			}
		} else {
			return fmt.Errorf("invalid argument")
		}
	}

	for _, arg := range c.Args() {
		if n, err := strconv.Atoi(arg); err == nil {
			pathList = PathList()
			if len(pathList) <= n {
				return fmt.Errorf("target index is not exists")
			}
			cleanDir(pathList[n]) // The path is indicated by number
		} else {
			cleanDir(arg) // Indicated by path name
		}
	}

	return nil
}
