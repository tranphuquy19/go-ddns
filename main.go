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

func main() {
	info()
	commands()

	// fmt.Println("value " + configPath)

	// test scheduler
	// scheduler.Every(5).Seconds().Do(task1)
	scheduler.StartAsync()

	err()

	for {
		select {}
	}
}
