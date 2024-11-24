package analyzer

import (
	"fmt"
	"io/ioutil"
	"os"
	"secure-lint/pkg/config"
	"secure-lint/pkg/models"
)

type Issue struct {
	Description string
	Severity    string
}

func gatherReport() string {
	outputDirectory := models.ProjectRoot + "/output/"
	files, err := os.ReadDir(outputDirectory)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return ""
	}

	var content string
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Skipping directory: %s\n", file.Name())
			continue
		}

		filePath := outputDirectory + file.Name()

		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
			continue
		}
		content += string(fileContent) + "\n"
	}
	return content
}

func outputReport() {
	fmt.Println(gatherReport())
}

func AnalyzeCode(path string, config *config.Config) {
	for _, analyzer := range config.Analyzers {

		if analyzer.CheckExecutable() {
			analyzer.Analyze(path)
		}
	}
	outputReport()
}
