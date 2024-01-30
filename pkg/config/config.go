package config

import (
	"bytes"
	"os"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v3"
)

// FromFile represent a config file content.
type FromFile struct {
	SourceRepositories []SourceRepository `yaml:"source_repositories"`
	TemplatePath       string             `yaml:"template_path"`
	OutputPath         string             `yaml:"output_path"`
}

// SourceRepository has details about a
// source repository to find ClusterApps in.
type SourceRepository struct {
	URL             string `yaml:"url"`
	Organization    string `yaml:"organization"`
	Name            string `yaml:"name"`
	CommitReference string `yaml:"commit_reference"`
}

// Read reads a config file and returns a struct.
func Read(path string) (*FromFile, error) {
	f := &FromFile{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadConfigFileError, err.Error())
	}

	reader := bytes.NewReader(data)
	decoder := yaml.NewDecoder(reader)
	// Fail on unknown fields.
	decoder.KnownFields(true)
	err = decoder.Decode(f)
	if err != nil {
		return nil, microerror.Maskf(CouldNotParseConfigFileError, err.Error())
	}

	return f, nil
}
