package main

import (
	"fmt"
	client "go-ddns/clients"
	"go-ddns/parser"
	"go-ddns/util"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

func runAction(c *cli.Context) error {
	context, _ := os.Getwd()
	if c.NArg() > 0 {
		a := c.Args().First()
		if a != "." {
			context = a
		}
	}
	if configPath == "ddns.yaml | ddns.json" {
		configPath = "ddns.yaml"
	}
	configFileFullPath := filepath.Join(context, configPath)
	configCredentialsPath := filepath.Join(context, credentialsPath)

	if !util.FileExists(configCredentialsPath) {
		forever = false
		fmt.Printf("Credentials file: %s does not exist\n", configCredentialsPath)
	}

	if util.FileExists(configFileFullPath) {
		config := parser.ConfigYAMLParser(configFileFullPath)
		forever = true
		for _, provider := range config.Providers {
			profile, providerName := provider.Profile, strings.ToLower(provider.Name)
			token := parser.TOMLGetProfile(configCredentialsPath, profile)
			fmt.Println("Using profile:", token)
			for _, domain := range provider.Domains {
				for _, record := range domain.Records {
					domainName := domain.Name
					recordType := strings.ToLower(record.Source.Type)
					triggerType, triggerValue := strings.ToLower(record.Trigger.Type), record.Trigger.Value
					switch recordType {
					case "get", "post":
						if triggerType == "cron_job" {
							switch providerName {
							case "netlify":
								scheduler.Cron(triggerValue).Do(func() {
									client.NetlifyUpdateRecord(domainName, &record, token)
								})
							}
						}
					}
				}
			}
		}
	} else {
		fmt.Printf("Config file: %s does not exist\n", configFileFullPath)
	}
	return nil
}

func profileAction(c *cli.Context) error {
	profileName := "default"
	if c.NArg() > 0 {
		profileName = c.Args().First()
	}
	fmt.Println(profileName)
	return nil
}

func info() {
	app.Name = "go-ddns"
	app.Usage = "a powerful dynamic DNS client tool in Go. Support: Cloudflare, Netlify, etc."
	app.Author = "github.com/tranphuquy19"
	app.Version = "0.0.2"
	app.EnableBashCompletion = true
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "profile",
			Aliases: []string{"p"},
			Usage:   "manage the profiles",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new profile",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "token, t",
							Usage:    "add API token (Clouldflare, Netlify, etc.): `TOKEN`",
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
			Action: profileAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "start the app",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Value:       "ddns.yaml | ddns.json",
					Required:    false,
					Usage:       "config file path. Supported YAML, JSON file",
					EnvVar:      "CONFIG_PATH",
					Destination: &configPath,
				},
				cli.StringFlag{
					Name:        "profile-path, p",
					Value:       ".credentials",
					Required:    false,
					Usage:       "profiles file path",
					EnvVar:      "CREDENTIALS_PATH",
					Destination: &credentialsPath,
				},
			},
			Action: runAction,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}

func err() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		forever = false
	}
}
