package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"secure-lint/pkg/analyzer"
	"secure-lint/pkg/config"
)

func AnalyzeCommand(config *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "analyze",
		Usage: "Analyze Go code for security vulnerabilities",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "path",
				Aliases:  []string{"p"},
				Usage:    "Path to the directory or file to analyze",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			path := c.String("path")
			fmt.Printf("Analyzing code at path: %s\n", path)
			analyzer.AnalyzeCode(path, config)
			return nil
		},
	}
}
