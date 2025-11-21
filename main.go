package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fyllekanin/murl/internal/commands"
	"github.com/fyllekanin/murl/internal/softwares"
)

func main() {
	includeUnstable := flag.Bool("unstable", false, "Include unstable versions")
	flag.Parse()

	argCount := len(flag.Args())
	items := make(map[string]softwares.Software)
	items["go"] = softwares.NewGoLang()
	items["nodejs"] = softwares.NewNodeJs()

	if argCount < 1 {
		fmt.Println("No arguments provided, try \"help\" to see usage")
		return
	}
	if strings.ToLower(flag.Args()[0]) == "help" {
		commands.Help()
		return
	}

	software, ok := items[strings.ToLower(flag.Args()[0])]
	if !ok {
		fmt.Println("Software " + flag.Args()[0] + " is not suported")
		return
	}

	switch flag.Args()[1] {
	case "list":
		commands.List(software, getBoolValue(includeUnstable))
	case "install":
		commands.Install(software, flag.Args()[2])
	default:
		fmt.Printf("Unknown command: \"%s\"\n", flag.Args()[1])
	}
}

func getBoolValue(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}
