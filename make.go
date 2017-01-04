package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

const configFileName = "config"

type item struct {
	name string
	dir  string
}

func (i *item) path() string {
	return filepath.Join(i.dir, i.name)
}

func parseArguments(c *cli.Context) (*item, error) {
	path := "."
	if c.NArg() == 1 {
		path = c.Args()[0]
	}

	name := c.String("name")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return nil, fmt.Errorf("cannot create directories: " + err.Error())
		}
	}

	_, err = os.Stat(filepath.Join(path, name))
	if err == nil {
		return nil, fmt.Errorf("target directory name " + filepath.Join(path, name) + " is exists")
	}

	return &item{name: name, dir: path}, nil
}

func register(path string) error {
	user, _ := user.Current()

	f, err := os.OpenFile(user.HomeDir+"/"+dirName+"/"+configFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("cannot register the directory: %s", err)
	}
	defer f.Close()

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("cannot get full path: %s", err)
	}
	f.WriteString(absPath + "\n")

	return nil
}

func makeDir(c *cli.Context) error {
	target, err := parseArguments(c)
	if err != nil {
		return fmt.Errorf("parseArguments error: %s", err)
	}

	err = os.Mkdir(target.path(), 0755)
	if err != nil {
		return fmt.Errorf("cannot create the temporary directory: %s", err)
	}

	register(target.path())

	fmt.Println(target.path() + " created.")

	return nil
}
