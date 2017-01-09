package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

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
	switch {
	case os.IsExist(err):
		return nil, fmt.Errorf("target directory name " + filepath.Join(path, name) + " is exists")
	case os.IsPermission(err):
		return nil, fmt.Errorf("permission denied")
	}

	return &item{name: name, dir: path}, nil
}

func register(path string) error {
	user, _ := user.Current()

	f, err := os.OpenFile(filepath.Join(user.HomeDir, dirName, configFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("cannot open the configuration directory: %s", err)
	}
	defer f.Close()

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("cannot get full path: %s", err)
	}
	f.WriteString(absPath + "\n")

	return nil
}

func unregister(path string) error {
	user, _ := user.Current()

	configPath := filepath.Join(user.HomeDir, dirName, configFileName)

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("cannot open the configuration directory: %s", err)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("cannot get full path: %s", err)
	}

	fmt.Println(absPath)
	ioutil.WriteFile(configPath, bytes.Replace(content, []byte(absPath), []byte(""), 1), 0644)

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
