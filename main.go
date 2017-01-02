package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/urfave/cli"
)

const dirName = ".tmp-cli"

func init() {
	user, _ := user.Current()
	tmpCLIPath := user.HomeDir + "/"

	_, err := os.Stat(tmpCLIPath + dirName)
	if os.IsNotExist(err) {
		err := os.Mkdir(tmpCLIPath+dirName, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot create directory "+tmpCLIPath+dirName+": "+err.Error())
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
