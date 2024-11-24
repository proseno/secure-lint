package report

import (
	"encoding/json"
	"os"
	"secure-lint/pkg/analyzer"
)

func GenerateJSONReport(issues []analyzer.Issue, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(issues)
}
