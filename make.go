package main

import (
	"fmt"
	"os"

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
	path := dir + "/"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0644)
		if err != nil {
			return fmt.Errorf("cannot create directories: " + path)
		}
	}

	_, err = os.Stat(path + name)
	if err == nil {
		return fmt.Errorf("target directory name " + path + name + " is exists")
	}

	err = os.Mkdir(path+name, 0644)
	if err != nil {
		return fmt.Errorf("cannot create the temporary directory: ", err)
	}

	fmt.Println("directory: " + dir + " name: " + name + " created.")

	return nil
}
