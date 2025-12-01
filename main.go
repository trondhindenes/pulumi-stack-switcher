package main

import "github.com/trond/pulumi-stack-switcher/cmd"

// Version is set at build time via ldflags
var Version = "dev"

func main() {
	cmd.SetVersion(Version)
	cmd.Execute()
}
