package main

import (
	"os"

	"github.com/rail44/hammer/cmd"
)

var version string

func main() {
	cmd.Version = version

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
