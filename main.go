package main

import (
	"github.com/thatisuday/commando"
)

// import (
// 	"go-ddns/parser"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	fileName := os.Args[1]
// 	workingDir, _ := os.Getwd()
// 	fullPath := filepath.Join(workingDir, fileName)
// 	parser.YAMLParser(fullPath)
// }

func main() {
	commando.
		SetExecutableName("go-ddns").
		SetVersion("0.0.1").
		SetDescription("A simple dynamic dns client in Go.")

	commando.
		Register(nil).
		AddFlag("config,c", "config file path. Supported YAML, JSON file (default: ddns.yaml or ddns.json)", commando.String, nil).
		AddFlag("profile-path,p", "profiles file path", commando.String, ".credentials").
		SetAction(handler)

	commando.
		Register("profile").
		SetShortDescription("manages the profile list").
		SetDescription("Description of profile").
		AddArgument("add", "create a new profile", "'default'").
		AddFlag("token,t", "API token (ClouldFlare, Netlify, etc.)", commando.String, "NONE").
		AddFlag("username,u", "your login username", commando.String, "NONE").
		AddFlag("password,p", "your login password", commando.String, "NONE").
		AddFlag("path,i", "path of credentials, where the profile file saved", commando.String, "./.credentials").
		AddArgument("remove", "remove existed profile", "'default'").
		SetAction(handler)

	commando.Parse(nil)
}
