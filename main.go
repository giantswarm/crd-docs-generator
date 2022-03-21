package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"

	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/crd-docs-generator/pkg/annotations"
	"github.com/giantswarm/crd-docs-generator/pkg/config"
	"github.com/giantswarm/crd-docs-generator/pkg/crd"
	"github.com/giantswarm/crd-docs-generator/pkg/git"
	"github.com/giantswarm/crd-docs-generator/pkg/output"
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

	// Default path for Markdown output (if not given in config)
	defaultOutputPath = "./output"
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

	// Full names and versions of CRDs found, to avoid duplicates.
	crdNameAndVersion := make(map[string]bool)

	outputPath := configuration.OutputPath
	if outputPath == "" {
		outputPath = defaultOutputPath
	}

	// Loop over configured repositories
	defer os.RemoveAll(repoFolder)
	for _, sourceRepo := range configuration.SourceRepositories {
		// List of source YAML files containing CRD definitions.
		crdFiles := make(map[string]bool)

		log.Printf("INFO - repo %s (%s)", sourceRepo.ShortName, sourceRepo.URL)
		clonePath := repoFolder + "/" + sourceRepo.Organization + "/" + sourceRepo.ShortName
		// Clone the repositories containing CRDs
		log.Printf("INFO - repo %s - cloning repository", sourceRepo.ShortName)
		err = git.CloneRepositoryShallow(
			sourceRepo.Organization,
			sourceRepo.ShortName,
			sourceRepo.CommitReference,
			clonePath)
		if err != nil {
			return microerror.Mask(err)
		}

		// Collect CRD YAML files
		for _, crdPath := range sourceRepo.CRDPaths {
			thisCRDFolder := clonePath + "/" + crdPath
			err = filepath.Walk(thisCRDFolder, func(path string, info os.FileInfo, err error) error {
				if strings.HasSuffix(path, ".yaml") {
					crdFiles[path] = true
				}
				return nil
			})
			if err != nil {
				return microerror.Mask(err)
			}
		}

		// Collect annotation info
		var repoAnnotations []annotations.CRDAnnotationSupport
		for _, annotationsPath := range sourceRepo.AnnotationsPath {
			thisAnnotationsFolder := clonePath + "/" + annotationsPath
			log.Printf("INFO - repo %s - collecting annotations in %s", sourceRepo.ShortName, thisAnnotationsFolder)
			a, err := annotations.Collect(thisAnnotationsFolder)
			if err != nil {
				log.Printf("ERROR - repo %s - collecting annotations in %s yielded error %#v", sourceRepo.ShortName, thisAnnotationsFolder, err)
			}
			repoAnnotations = append(repoAnnotations, a...)
		}

		crdFilesSlice := []string{}
		for crdFile := range crdFiles {
			crdFilesSlice = append(crdFilesSlice, crdFile)
		}

		sort.Strings(crdFilesSlice)
		for _, crdFile := range crdFilesSlice {
			log.Printf("INFO - repo %s - reading CRDs from file %s", sourceRepo.ShortName, crdFile)

			crds, err := crd.Read(crdFile)
			if err != nil {
				log.Printf("WARN - something went wrong in crd.Read for file %s: %#v", crdFile, err)
			}

			for i := range crds {
				versions := []string{}

				for _, v := range crds[i].Spec.Versions {
					fullKey := crds[i].Name + "_" + v.Name

					_, exists := crdNameAndVersion[fullKey]
					if exists {
						log.Printf("WARN - repo %s - provides CRD %s version %s which is already added - skipping", sourceRepo.ShortName, crds[i].Name, v.Name)
						continue
					}
					crdNameAndVersion[fullKey] = true
					versions = append(versions, v.Name)
				}

				log.Printf("INFO - repo %s - processing CRD %s with versions %s", sourceRepo.ShortName, crds[i].Name, versions)

				// Skip hidden CRDs and CRDs with missing metadata
				meta, ok := sourceRepo.Metadata[crds[i].Name]
				if !ok {
					log.Printf("WARN - repo %s - skipping %s as no metadata found", sourceRepo.ShortName, crds[i].Name)
					continue
				}
				if meta.Hidden {
					log.Printf("INFO - repo %s - skipping %s as hidden by configuration", sourceRepo.ShortName, crds[i].Name)
					continue
				}

				// Get at most one example CR for each version of this CRD
				exampleCRs := make(map[string]string)
				for _, version := range versions {
					found := false

					for _, crPath := range sourceRepo.CRPaths {

						crFilePath := fmt.Sprintf("%s/%s/%s_%s_%s.yaml", clonePath, crPath, crds[i].Spec.Group, version, crds[i].Spec.Names.Singular)
						if _, err := os.Stat(crFilePath); errors.Is(err, os.ErrNotExist) {
							continue
						}

						exampleCR, err := ioutil.ReadFile(crFilePath)
						if err != nil {
							log.Printf("ERROR - repo %s - example CR %s could not be read: %s", sourceRepo.ShortName, crFilePath, err)
						} else {
							found = true
							exampleCRs[version] = strings.TrimSpace(string(exampleCR))
							break
						}
					}

					if !found {
						log.Printf("WARN - repo %s - No example CR found for %s version %s ", sourceRepo.ShortName, crds[i].Name, version)
					}
				}

				templatePath := path.Dir(configFilePath) + "/" + configuration.TemplatePath

				crdAnnotations := annotations.FilterForCRD(repoAnnotations, crds[i].Name, "")

				_, err = output.WritePage(
					crds[i],
					crdAnnotations,
					meta,
					exampleCRs,
					outputPath,
					sourceRepo.URL,
					sourceRepo.CommitReference,
					templatePath)
				if err != nil {
					log.Printf("WARN - repo %s - something went wrong in WriteCRDDocs: %#v", sourceRepo.ShortName, err)
				}
			}
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
