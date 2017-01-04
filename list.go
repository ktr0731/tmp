package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func list(c *cli.Context) error {
	user, _ := user.Current()

	content, err := ioutil.ReadFile(filepath.Join(user.HomeDir, dirName, configFileName))
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(content), "\n") {
		fmt.Println(line)
	}

	return nil
}
