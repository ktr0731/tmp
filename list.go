package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func list(c *cli.Context) error {
	index := 0
	list := PathList()
	for _, item := range list {
		if c.BoolT("number") {
			fmt.Printf("[%d] %s", index, item)
		}
		fmt.Println(item)
	}

	return nil
}
