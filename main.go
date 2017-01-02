package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const dirName = ".tmp-cli"

func init() {
	_, err := os.Stat("~/" + dirName)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot create directory "+dirName+": "+err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "tmp"
	app.Version = "0.1.0"
	app.Usage = "Easily manage your all temporary directories"
	app.Commands = commands
	app.Run(os.Args)
}
