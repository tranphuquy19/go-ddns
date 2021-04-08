package main

import (
	"github.com/urfave/cli"
)

var app = cli.NewApp()

var configPath string
var profilePath string

func main() {
	info()
	commands()

	// fmt.Println("value " + configPath)

	err()
}
