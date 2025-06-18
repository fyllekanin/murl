package commands

import "fmt"

func Help() {
	fmt.Println(`Murl is a tool to help manage software versions available on the system.

Usage:
    murl [flags] <software> <command>

Available softwares:
    go
    nodejs
    maven
    java

Available commands:
    list       lists the available versions
        -unstable to include unstable versions if applicable`)
}
