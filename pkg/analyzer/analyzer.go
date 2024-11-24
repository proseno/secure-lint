package analyzer

import (
	"fmt"
	"secure-lint/pkg/config"
)

type Issue struct {
	Description string
	Severity    string
}

func AnalyzeCode(path string, config *config.Config) {
	for _, analyzer := range config.Analyzers {

		if analyzer.CheckExecutable() {
			analyzeResult := analyzer.Analyze(path)
			fmt.Println(analyzeResult)
		}
	}
}
