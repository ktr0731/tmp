package main

import (
	"bytes"
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

	for _, test := range tests {
		args := strings.Split(test.command, " ")
		app := cli.NewApp()
		app.Writer, app.ErrWriter = outStream, errStream

		app.Run(args)

		// context := &cli.Context{App: app, Command: *(app.Command("make"))}
		//
		// _, err := parseArguments(context)
		// if err != nil && test.success {
		// 	t.Error(err)
		// }
	}
}
