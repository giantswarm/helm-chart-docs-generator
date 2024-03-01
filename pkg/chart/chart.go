package chart

import (
	"log"
	"os"
	"os/exec"

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
func GenerateChartConfig(basePath string, chartName string) ([]byte, error) {
	cmd := exec.Command("go", "install", "github.com/giantswarm/schemadocs@latest")
	err := cmd.Run()
	if err != nil {
		return nil, microerror.Maskf(CouldNotGenerateChartFileError, err.Error())
	}
	cmd = exec.Command("schemadocs", "generate", "values.schema.json", "-o", "README.md", "-l", "linear")
	cmd.Dir = basePath + "/helm/" + chartName
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, microerror.Maskf(CouldNotGenerateChartFileError, err.Error(), string(output))
	}
	content, err := os.ReadFile(basePath + "/helm/" + chartName + "/README.md")
	if err != nil {
		return nil, microerror.Maskf(CouldNotGenerateChartFileError, err.Error())
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
		return m, microerror.Maskf(CouldNotGenerateChartFileError, err.Error())
	}

	err = yaml.Unmarshal(metadata, &m)
	if err != nil {
		return m, microerror.Maskf(CouldNotParsedChartFileError, err.Error())
	}

	return m, nil
}
