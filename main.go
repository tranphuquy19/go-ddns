package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/urfave/cli"
)

var app = cli.NewApp()
var scheduler = gocron.NewScheduler(time.UTC)

var configPath string
var profilePath string
var forever bool = false

func main() {
	info()
	commands()

	scheduler.StartAsync()

	err()

	if forever {
		for {
			select {}
		}
	}
}
