package chart

import (
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

const HELM_CHARTS_FOLDER = "/helm/"

type Metadata struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Annotations struct {
		Team string `yaml:"application.giantswarm.io/team"`
	} `yaml:"annotations"`
}

// GenerateChartConfig generates a README YAML file and returns the Content to render.
func GenerateChartConfig(basePath string, chartName string) ([]byte, error) {
	cmd := exec.Command("schemadocs", "generate", "values.schema.json", "-o", "README.md", "-l", "linear")
	cmd.Dir = basePath + HELM_CHARTS_FOLDER + chartName
	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(basePath + HELM_CHARTS_FOLDER + chartName + "/README.md")
	if err != nil {
		return nil, err
	}

	return content, nil
}

// Read reads a README YAML file and returns the Content to render.
func ReadChartMetadata(basePath string, chartName string) (Metadata, error) {
	var m Metadata
	chartPath := basePath + HELM_CHARTS_FOLDER + chartName + "/Chart.yaml"

	log.Printf("INFO - chart %s - reading Chart yaml", chartPath)
	metadata, err := os.ReadFile(chartPath)
	if err != nil {
		return m, err
	}

	err = yaml.Unmarshal(metadata, &m)
	if err != nil {
		return m, err
	}

	return m, nil
}
