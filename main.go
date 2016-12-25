package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tmp"
	app.Version = "0.1.0"
	app.Usage = "Easily manage your all temporary directories"
	app.Commands = commands
	app.Run(os.Args)
}
