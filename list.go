package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

func list(c *cli.Context) error {
	user, _ := user.Current()

	f, err := os.Open(filepath.Join(user.HomeDir, dirName, configFileName))
	if err != nil {
		return fmt.Errorf("cannot open tmp config file: %s", err.Error())
	}

	list := ""
	index := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if c.BoolT("number") {
			line = fmt.Sprintf("[%d] ", index) + line
		}
		line += "\n"

		list += line

		index++
	}

	fmt.Println(list)

	return nil
}
