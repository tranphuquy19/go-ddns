package main

import (
	"fmt"

	"github.com/thatisuday/commando"
)

func handler(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	fmt.Printf("Printing options of the `root` command...\n\n")

	// print arguments
	for k, v := range args {
		fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
	}

	// print flags
	for k, v := range flags {
		fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
	}
}
