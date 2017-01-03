package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/urfave/cli"
)

func TestParseArguments(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)

	tests := []struct {
		command string
		success bool
	}{
		{"tmp make", true},
	}

	app := cli.NewApp()
	app.Writer, app.ErrWriter = outStream, errStream
	app.Commands = commands

	for _, test := range tests {
		args := strings.Split(test.command, " ")
		if err := app.Run(args); err != nil {
			t.Fatal(err)
		}

		os.Remove("./tmp/")
	}
}
