package config

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// FromFile represent a config file content.
type FromFile struct {
	SourceRepositories []SourceRepository `yaml:"source_repositories"`
	TemplatePath       string             `yaml:"template_path"`
	OutputPath         string             `yaml:"output_path"`
}

// SourceRepository has details about a
// source repository to find HelmCharts in.
type SourceRepository struct {
	URL             string `yaml:"url"`
	Organization    string `yaml:"organization"`
	Introduction    string `yaml:"introduction"`
	Name            string `yaml:"name"`
	CommitReference string `yaml:"commit_reference"`
}

// Read reads a config file and returns a struct.
func Read(path string) (*FromFile, error) {
	f := &FromFile{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file %q with %w", path, err)
	}

	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	// Fail on unknown fields.
	decoder.KnownFields(true)
	err = decoder.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the file content %q with %w", path, err)
	}

	return f, nil
}
