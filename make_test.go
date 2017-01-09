package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"

	"github.com/urfave/cli"
)

func TestMakeDir(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	tests := []struct {
		command string
		path    string
		name    string
		success bool
	}{
		{
			command: "tmp make temp",
			success: true,
		},
	}

	app := cli.NewApp()
	app.Writer, app.ErrWriter = outStream, errStream
	app.Commands = commands

	user, _ := user.Current()

	for _, test := range tests {
		args := strings.Split(test.command, " ")
		if err := app.Run(args); err != nil {
			t.Fatal(err)
		}

		if test.name == "" {
			test.name = "tmp"
		}

		os.Remove(filepath.Join(test.path, test.name))

		// Tear down
		configFilePath := filepath.Join(user.HomeDir, dirName, configFileName)
		contentBytes, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			t.Fatal(err)
		}
		content := string(contentBytes)
		splitted := strings.Split(content, "\n")

		if content != "" {
			// Remove item from ~/.tmp-cli/config
			ioutil.WriteFile(configFilePath, []byte(strings.Join(splitted[:len(splitted)-2], "\n")), 0644)
		}
	}
}
