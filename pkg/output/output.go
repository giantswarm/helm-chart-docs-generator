package output

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/helm-chart-docs-generator/pkg/chart"
)

// PageData is all the data we pass to generate the Chart detail page.
type PageData struct {
	Description         string
	SourceRepository    string
	SourceRepositoryRef string
	Title               string
	Team                string
	Introduction        string
	Content             string
	Weight              int
}

// WritePage creates a Cluster App configuration page.
func WritePage(
	metadata chart.Metadata,
	content,
	introduction,
	outputFolder,
	repoURL,
	repoRef,
	templatePath string) (string, error) {

	templateCode, err := os.ReadFile(templatePath)
	if err != nil {
		return "", microerror.Maskf(cannotOpenTemplate, "Could not read template file %s: %s", templatePath, err)
	}

	// Add custom functions support for our template.
	funcMap := sprig.FuncMap()
	// Join strings by separator
	funcMap["join"] = strings.Join
	funcMap["currentDate"] = CurrentDate

	// Read our output template.
	tpl := template.Must(template.New("schemapage").Funcs(funcMap).Parse(string(templateCode)))

	// Collect values to pass to our output template.
	data := PageData{
		SourceRepository:    repoURL,
		SourceRepositoryRef: repoRef,
		Title:               metadata.Name,
		Description:         metadata.Description,
		Team:                metadata.Annotations.Team,
		Weight:              100,
		Introduction:        introduction,
		Content:             content,
	}

	// Name output file after full Cluster App name.
	outputFile := outputFolder + "/" + metadata.Name + ".md"

	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.MkdirAll(outputFolder, os.ModePerm)
		if err != nil {
			return "", microerror.Mask(err)
		}
	}

	handler, err := os.Create(outputFile)
	if err != nil {
		return "", microerror.Mask(err)
	}

	err = tpl.Execute(handler, data)
	if err != nil {
		fmt.Printf("%s: %s\n", outputFile, err)
	}

	return outputFile, nil
}

func CurrentDate(format string) string {
	return time.Now().Format(format)
}
