package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"secure-lint/pkg/analyzer"
	"secure-lint/pkg/config"
	"secure-lint/pkg/models"
	"strings"
)

func getPath(c *cli.Context) string {
	return c.String("path")
}

func getLangs(c *cli.Context) []string {
	langString := c.String("lang")
	return strings.Split(langString, ",")
}

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
			&cli.StringFlag{
				Name:     "lang",
				Aliases:  []string{"l"},
				Usage:    "Comma separated languages to analyze",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			path := getPath(c)
			langs := getLangs(c)
			path = models.ProjectRoot + "/" + path
			fmt.Printf("Analyzing code at path: %s\n", path)
			analyzer.AnalyzeCode(path, config, langs)
			return nil
		},
	}
}
