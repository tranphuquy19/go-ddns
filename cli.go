package main

import "github.com/urfave/cli"

func info() {
	app.Name = "go-ddns"
	app.Usage = "A simple dynamic dns client in Go."
	app.Author = "github.com/tranphuquy19"
	app.Version = "0.0.2"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "profile",
			Usage: "manage the profiles",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new profile",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "token, t",
							Usage:    "add API token (ClouldFlare, Netlify, etc.): `TOKEN`",
							Required: false,
						},
						cli.StringFlag{
							Name:     "username, u",
							Usage:    "add username: `USERNAME`",
							Required: false,
						},
						cli.StringFlag{
							Name:     "password, p",
							Usage:    "add password: `PASSWORD`",
							Required: false,
						},
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing profile",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "name, n",
							Usage:    "remove an existing profile by name",
							Required: true,
						},
					},
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config, c",
			Value:       "ddns.yaml | ddns.json",
			Usage:       "config file path. Supported YAML, JSON file",
			Destination: &configPath,
		},
		cli.StringFlag{
			Name:        "profile-path, p",
			Value:       ".credentials",
			Usage:       "profiles file path",
			Destination: &profilePath,
		},
	}
}
