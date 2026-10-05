package main

import (
	"github.com/CamilleGR/goplant/internal/cli"
	"os"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
