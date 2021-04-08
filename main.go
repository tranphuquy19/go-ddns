package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var configPath string
var profilePath string

func main() {
	info()
	commands()

	fmt.Println("value " + configPath)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
