package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func list(c *cli.Context) error {
	// user, _ := user.Current()
	//
	// f, err := os.Open(filepath.Join(user.HomeDir, dirName, configFileName))
	// if err != nil {
	// 	return fmt.Errorf("cannot open tmp config file: %s", err.Error())
	// }
	//
	// list := ""
	// index := 0
	// s := bufio.NewScanner(f)
	// for s.Scan() {
	// 	line := s.Text()
	// 	if c.BoolT("number") {
	// 		line = fmt.Sprintf("[%d] ", index) + line
	// 	}
	// 	line += "\n"
	//
	// 	list += line
	//
	// 	index++
	// }
	//
	// fmt.Println(list)

	list := PathList()
	for _, item := range list {
		fmt.Println(item)
	}

	return nil
}
