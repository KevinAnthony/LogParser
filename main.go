package main

import (
	"fmt"
	"os"

	"github.com/KevinAnthony/LogParser/sealer"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Version: "v0.0.1",
		Usage:   "Switch between clusters.",
		Commands: []*cli.Command{
			sealer.Parse(),
		},
		Flags:                  nil,
		EnableBashCompletion:   true,
		UseShortOptionHandling: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // nolint: gomnd
	}
}
