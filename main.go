package main

import (
	"log"
	"os"
	"secure-lint/pkg/config"

	"github.com/urfave/cli/v2"
	"secure-lint/cmd"
)

func main() {

	configData, _ := config.LoadConfig("config.yaml")
	app := &cli.App{
		Name:  "secure-lint",
		Usage: "A linting tool to analyze Go code for security vulnerabilities.",
		Commands: []*cli.Command{
			cmd.AnalyzeCommand(configData),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
