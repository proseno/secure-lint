package config

import (
	"fmt"
	"secure-lint/pkg/models"

	"gopkg.in/yaml.v3"

	"os"
)

type Config struct {
	Analyzers []models.Analyzer `yaml:"analyzers"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &config, nil
}

func (c *Config) GetByLang(lang string) (*models.Analyzer, error) {
	for _, analyzer := range c.Analyzers {
		if analyzer.As == lang {
			return &analyzer, nil
		}
	}
	return nil, fmt.Errorf("analyzer for language %s not found in config", lang)
}
