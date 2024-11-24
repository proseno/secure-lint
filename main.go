package main

import (
	"log"
	"os"
	"secure-lint/pkg/config"
	"secure-lint/pkg/models"

	"github.com/urfave/cli/v2"
	"secure-lint/cmd"
)

func initProjectRoot() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	models.ProjectRoot = projectRoot
}

func main() {
	configData, _ := config.LoadConfig("config.yaml")
	initProjectRoot()
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
