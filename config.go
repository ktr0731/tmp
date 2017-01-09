package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	dirName        = ".tmp-cli"
	configFileName = "config"
)

var pathList []string

func loadConfig() error {
	user, _ := user.Current()

	f, err := os.Open(filepath.Join(user.HomeDir, dirName, configFileName))
	if err != nil {
		return fmt.Errorf("cannot open tmp config file: %s", err.Error())
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		pathList = append(pathList, string(s.Text()))
	}

	return nil
}

func contains(target string) bool {
	for _, path := range PathList() {
		if path == target {
			return true
		}
	}

	return false
}

func PathList() []string {
	loadConfig()

	return pathList
}
