package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/crd-docs-generator/service/config"
	"github.com/giantswarm/crd-docs-generator/service/crd"
	"github.com/giantswarm/crd-docs-generator/service/git"
	"github.com/giantswarm/crd-docs-generator/service/output"
)

// CRDDocsGenerator represents an instance of this command line tool, it carries
// the cobra command which runs the process along with configuration parameters
// which come in as flags on the command line.
type CRDDocsGenerator struct {
	// Internals.
	rootCommand *cobra.Command

	// Settings/Preferences

	// Path to the config file
	configFilePath string
}

const (
	// Target path for our clone of the apiextensions repo.
	repoFolder = "/tmp/gitclone"

	crdFolder = repoFolder + "/config/crd/v1"

	crFolder = repoFolder + "/docs/cr"

	// Path for Markdown output.
	outputFolderPath = "./output"

	// Path for templates
	templateFolderPath = "./templates"

	// Single CRD page template filename (without path)
	outputTemplate = "crd.template"
)

func main() {
	var err error

	var crdDocsGenerator CRDDocsGenerator
	{
		c := &cobra.Command{
			Use:          "crd-docs-generator",
			Short:        "crd-docs-generator is a command line tool for generating markdown files that document Giant Swarm's custom resources",
			SilenceUsage: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return generateCrdDocs(crdDocsGenerator.configFilePath)
			},
		}

		c.PersistentFlags().StringVar(&crdDocsGenerator.configFilePath, "config", "./config.yaml", "Path to the configuration file.")
		crdDocsGenerator.rootCommand = c
	}

	if err = crdDocsGenerator.rootCommand.Execute(); err != nil {
		printStackTrace(err)
		os.Exit(1)
	}
}

// generateCrdDocs is the function called from our main CLI command.
func generateCrdDocs(configFilePath string) error {
	configuration, err := config.Read(configFilePath)
	if err != nil {
		return microerror.Mask(err)
	}

	crdFiles := []string{}

	err = git.CloneRepositoryShallow(
		configuration.SourceRepository.Organization,
		configuration.SourceRepository.ShortName,
		configuration.SourceRepository.CommitReference,
		repoFolder)
	if err != nil {
		return microerror.Mask(err)
	}

	defer os.RemoveAll(repoFolder)

	err = filepath.Walk(crdFolder, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if strings.HasSuffix(path, ".yaml") {
			fmt.Printf("Collecting file %s\n", path)
			crdFiles = append(crdFiles, path)
		}
		return nil
	})
	if err != nil {
		return microerror.Mask(err)
	}

	for _, crdFile := range crdFiles {
		fmt.Printf("Reading file %s\n", crdFile)

		crd, err := crd.Read(crdFile)
		if err != nil {
			fmt.Printf("Something went wrong in ReadCRD: %#v\n", err)
		}

		if contains(configuration.SkipCRDs, crd.Name) {
			fmt.Printf("Skipping CRD %s\n", crd.Name)
		} else {
			fmt.Printf("Writing output for CRD %s\n", crd.Name)

			err = output.WritePage(
				crd,
				crFolder,
				outputFolderPath,
				configuration.SourceRepository.URL,
				configuration.SourceRepository.CommitReference,
				templateFolderPath,
				outputTemplate)
			if err != nil {
				fmt.Printf("Something went wrong in WriteCRDDocs: %#v\n", err)
			}
		}
	}

	return nil
}

// contains checks whether slice contains the given item.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func printStackTrace(err error) {
	fmt.Println("\n--- Stack Trace ---")
	var stackedError microerror.JSONError
	jsonErr := json.Unmarshal([]byte(microerror.JSON(err)), &stackedError)
	if jsonErr != nil {
		fmt.Println("Error when trying to Unmarshal JSON error:")
		fmt.Printf("%#v\n", jsonErr)
		fmt.Println("\nOriginal error:")
		fmt.Printf("%#v\n", err)
	}

	for i, j := 0, len(stackedError.Stack)-1; i < j; i, j = i+1, j-1 {
		stackedError.Stack[i], stackedError.Stack[j] = stackedError.Stack[j], stackedError.Stack[i]
	}

	for _, entry := range stackedError.Stack {
		fmt.Printf("%s:%d\n", entry.File, entry.Line)
	}
}
