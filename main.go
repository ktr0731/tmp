package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/urfave/cli"
)

func init() {
	user, _ := user.Current()

	path := filepath.Join(user.HomeDir, dirName)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot create directory "+path+": "+err.Error())
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
