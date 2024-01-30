package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/cluster-app-docs-generator/pkg/chart"
	"github.com/giantswarm/cluster-app-docs-generator/pkg/config"
	"github.com/giantswarm/cluster-app-docs-generator/pkg/git"
	"github.com/giantswarm/cluster-app-docs-generator/pkg/output"
)

// ClusterAppDocsGenerator represents an instance of this command line tool, it carries
// the cobra command which runs the process along with configuration parameters
// which come in as flags on the command line.
type ClusterAppDocsGenerator struct {
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

	var clusterAppDocsGenerator ClusterAppDocsGenerator
	{
		c := &cobra.Command{
			Use:          "cluster-app-docs-generator",
			Short:        "cluster-app-docs-generator is a command line tool for generating markdown files that document Giant Swarm's Cluster Apps.",
			SilenceUsage: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return generateClusterAppDocs(clusterAppDocsGenerator.configFilePath)
			},
		}

		c.PersistentFlags().StringVar(&clusterAppDocsGenerator.configFilePath, "config", "./config.yaml", "Path to the configuration file.")
		clusterAppDocsGenerator.rootCommand = c
	}

	if err = clusterAppDocsGenerator.rootCommand.Execute(); err != nil {
		printStackTrace(err)
		os.Exit(1)
	}
}

// generateClusterAppDocs is the function called from our main CLI command.
func generateClusterAppDocs(configFilePath string) error {
	configuration, err := config.Read(configFilePath)
	if err != nil {
		return microerror.Mask(err)
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
			return microerror.Mask(err)
		}

		log.Printf("INFO - repo %s - repository cloned successfully", sourceRepo.Name)
		// Collect Chart README YAML files
		chartMetadata, err := chart.ReadChartMetadata(clonePath, sourceRepo.Name)
		if err != nil {
			log.Fatal(err)
		}
		chartContent, err := chart.ReadChartConfig(clonePath, sourceRepo.Name)
		if err != nil {
			log.Fatal(err)
		}

		templatePath := path.Dir(configFilePath) + "/" + configuration.TemplatePath

		_, err = output.WritePage(
			chartMetadata,
			string(chartContent),
			outputPath,
			sourceRepo.URL,
			sourceRepo.CommitReference,
			templatePath)
		if err != nil {
			log.Printf("WARN - repo %s - something went wrong in WriteClusterAppDocs: %#v", sourceRepo.Name, err)
		}
	}

	return nil
}

func printStackTrace(err error) {
	fmt.Println("\n--- Stack Trace ---")
	var stackedError microerror.JSONError
	jsonErr := json.Unmarshal([]byte(microerror.JSON(err)), &stackedError)
	if jsonErr != nil {
		fmt.Println("Error when trying to Unmarshal JSON error:")
		log.Printf("%#v", jsonErr)
		fmt.Println("\nOriginal error:")
		log.Printf("%#v", err)
	}

	for i, j := 0, len(stackedError.Stack)-1; i < j; i, j = i+1, j-1 {
		stackedError.Stack[i], stackedError.Stack[j] = stackedError.Stack[j], stackedError.Stack[i]
	}

	for _, entry := range stackedError.Stack {
		log.Printf("%s:%d", entry.File, entry.Line)
	}
}
