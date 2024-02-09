package chart

import (
	"log"
	"os"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v3"
)

type Metadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Annotations struct {
		Team string `yaml:"application.giantswarm.io/team"`
	} `yaml:"annotations"`
}

// Read reads a README YAML file and returns the Content to render.
func ReadChartConfig(basePath string, chartName string) ([]byte, error) {
	content, err := os.ReadFile(basePath + "/helm/" + chartName + "/README.md")
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadChartFileError, err.Error())
	}

	return content, nil
}

// Read reads a README YAML file and returns the Content to render.
func ReadChartMetadata(basePath string, chartName string) (Metadata, error) {
	var m Metadata
	chartPath := basePath + "/helm/" + chartName + "/Chart.yaml"

	log.Printf("INFO - chart %s - reading Chart yaml", chartPath)
	metadata, err := os.ReadFile(chartPath)
	if err != nil {
		return m, microerror.Maskf(CouldNotReadChartFileError, err.Error())
	}

	err = yaml.Unmarshal(metadata, &m)
	if err != nil {
		return m, microerror.Maskf(CouldNotParsedChartFileError, err.Error())
	}

	return m, nil
}
