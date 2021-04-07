package main

import (
	"fmt"

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
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `root` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	commando.Parse(nil)
}
