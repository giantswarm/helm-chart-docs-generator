package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/giantswarm/helm-chart-docs-generator/pkg/chart"
	"github.com/giantswarm/helm-chart-docs-generator/pkg/config"
	"github.com/giantswarm/helm-chart-docs-generator/pkg/git"
	"github.com/giantswarm/helm-chart-docs-generator/pkg/output"
)

// helmChartDocsGenerator represents an instance of this command line tool, it carries
// the cobra command which runs the process along with configuration parameters
// which come in as flags on the command line.
type helmChartDocsGenerator struct {
	// Internals.
	rootCommand *cobra.Command

	// Settings/Preferences

	// Path to the config file
	configFilePath string
}

const (
	// Target path for our clone of the cluster apps repos.
	repoFolder = "/tmp/gitclone"

	// Default path for Markdown output (if not given in config)
	defaultOutputPath = "./output"
)

func main() {
	var err error

	var helmChartDocsGenerator helmChartDocsGenerator
	{
		c := &cobra.Command{
			Use:          "helm-chart-docs-generator",
			Short:        "helm-chart-docs-generator is a command line tool for generating markdown files that document Giant Swarm's Cluster Apps.",
			SilenceUsage: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return generateHelmChartDocs(helmChartDocsGenerator.configFilePath)
			},
		}

		c.PersistentFlags().StringVar(&helmChartDocsGenerator.configFilePath, "config", "./config.yaml", "Path to the configuration file.")
		helmChartDocsGenerator.rootCommand = c
	}

	if err = helmChartDocsGenerator.rootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// generateHelmChartDocs is the function called from our main CLI command.
func generateHelmChartDocs(configFilePath string) error {
	configuration, err := config.Read(configFilePath)
	if err != nil {
		return err
	}

	outputPath := configuration.OutputPath
	if outputPath == "" {
		outputPath = defaultOutputPath
	}

	// Loop over configured repositories
	defer os.RemoveAll(repoFolder)
	for _, sourceRepo := range configuration.SourceRepositories {
		log.Printf("INFO - repo %s (%s)", sourceRepo.Name, sourceRepo.URL)
		clonePath := repoFolder + "/" + sourceRepo.Organization + "/" + sourceRepo.Name
		// Clone the repositories containing Cluster Apps
		log.Printf("INFO - repo %s - cloning repository", sourceRepo.Name)
		err = git.CloneRepositoryShallow(
			sourceRepo.Organization,
			sourceRepo.Name,
			sourceRepo.CommitReference,
			clonePath)
		if err != nil {
			return err
		}

		log.Printf("INFO - repo %s - repository cloned successfully", sourceRepo.Name)
		// Collect Chart README YAML files
		chartMetadata, err := chart.ReadChartMetadata(clonePath, sourceRepo.Name)
		if err != nil {
			log.Fatal(err)
		}
		chartContent, err := chart.GenerateChartConfig(clonePath, sourceRepo.Name)
		if err != nil {
			log.Fatal(err)
		}

		templatePath := path.Dir(configFilePath) + "/" + configuration.TemplatePath

		_, err = output.WritePage(
			chartMetadata,
			string(chartContent),
			sourceRepo.Introduction,
			outputPath,
			sourceRepo.URL,
			sourceRepo.CommitReference,
			templatePath)
		if err != nil {
			log.Printf("WARN - repo %s - something went wrong in WriteHelmChartDocs: %#v", sourceRepo.Name, err)
		}
	}

	return nil
}
