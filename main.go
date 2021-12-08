package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

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

// Types to read annoatations documentation and CRD support
type AnnotationSupportRelease struct {
	Release    string
	APIVersion string
	CRD        string
}

type Annotation struct {
	Documentation string
	Support       []AnnotationSupportRelease
}

const (
	// Target path for our clone of the apiextensions repo.
	repoFolder = "/tmp/gitclone"

	// Within a git clone, relative path for Giant Swarm CRDs in YAML format.
	crdFolder = "config/crd"

	// Within a git clone, relative path for upstream CRDs in YAML format.
	upstreamCRDFolder = "helm"

	// File name for bespoke upstream CRDs.
	upstreamFileName = "upstream.yaml"

	// Within a git clone, relative path for example CRs in YAML format.
	crFolder = "docs/cr"

	annotationsFolder = "/pkg/annotation"

	// Path for Markdown output.
	outputFolderPath = "./output"
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

	// Full names of CRDs found
	crdNames := make(map[string]bool)

	// Loop over configured repositories
	defer os.RemoveAll(repoFolder)
	for _, sourceRepo := range configuration.SourceRepositories {
		// List of source YAML files containing CRD definitions.
		crdFiles := make(map[string]bool)

		// List of Go files containing annotation definitions.
		annotationFiles := []string{}

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

		// Collect annotation info
		thisAnnotationsFolder := clonePath + "/" + annotationsFolder
		log.Printf("INFO - repo %s - collecting annotations in %s", sourceRepo.ShortName, thisAnnotationsFolder)
		err = filepath.Walk(thisAnnotationsFolder, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".go") {
				annotationFiles = append(annotationFiles, path)
			}
			return nil
		})
		if err != nil {
			return microerror.Mask(err)
		}

		// Collect our own CRD YAML files
		thisCRDFolder := clonePath + "/" + crdFolder
		err = filepath.Walk(thisCRDFolder, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".yaml") {
				crdFiles[path] = true
			}
			return nil
		})
		if err != nil {
			return microerror.Mask(err)
		}

		// Collect upstream CRD YAML files
		thisUpstreamCRDFolder := clonePath + "/" + upstreamCRDFolder
		err = filepath.Walk(thisUpstreamCRDFolder, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, upstreamFileName) {
				crdFiles[path] = true
			}
			return nil
		})
		if err != nil {
			return microerror.Mask(err)
		}

		// Process annotation details
		var annotations []output.CRDAnnotationSupport
		for _, annotationFile := range annotationFiles {
			fset := token.NewFileSet()
			files := []*ast.File{
				mustParse(fset, annotationFile),
			}
			p, err := doc.NewFromFiles(fset, files, "github.com/giantswarm/extract-go-doc/p")
			if err != nil {
				panic(err)
			}
			for _, constant := range p.Consts {
				annotation := Annotation{}

				err := yaml.Unmarshal([]byte(constant.Doc), &annotation)
				if err != nil {
					log.Printf("%s - %s - annotation YAML docs missing", annotationFile, constant.Names[0])
				}

				if annotation.Documentation != "" {
					for _, crdAPI := range annotation.Support {

						rawAnnotation := constant.Decl.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value
						annotationValue := strings.Replace(rawAnnotation, "\"", "", -1)

						annotations = append(annotations, output.CRDAnnotationSupport{
							Annotation:    annotationValue,
							APIVersion:    crdAPI.APIVersion,
							CRDName:       crdAPI.CRD,
							Release:       crdAPI.Release,
							Documentation: annotation.Documentation,
						})
					}
				}
			}
		}

		for crdFile := range crdFiles {
			log.Printf("INFO - repo %s - reading CRDs from file %s", sourceRepo.ShortName, crdFile)

			crds, err := crd.Read(crdFile)
			if err != nil {
				log.Printf("WARN - something went wrong in crd.Read for file %s: %#v", crdFile, err)
			}

			for i := range crds {
				versions := []string{}
				for _, v := range crds[i].Spec.Versions {
					versions = append(versions, v.Name)
				}
				log.Printf("INFO - repo %s - processing CRD %s with versions %s", sourceRepo.ShortName, crds[i].Name, versions)

				_, exists := crdNames[crds[i].Name]
				if exists {
					log.Printf("WARN - repo %s - provides CRD %s which is already added - skipping", sourceRepo.ShortName, crds[i].Name)
					continue
				}
				crdNames[crds[i].Name] = true

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

				// Get example CRs for this CRD (using version as key)
				exampleCRs := make(map[string]string)
				for _, version := range versions {
					crFileName := fmt.Sprintf("%s/%s/%s_%s_%s.yaml", clonePath, crFolder, crds[i].Spec.Group, version, crds[i].Spec.Names.Singular)
					exampleCR, err := ioutil.ReadFile(crFileName)
					if err != nil {
						log.Printf("WARN - repo %s - CR example is missing for %s version %s in path %s", sourceRepo.ShortName, crds[i].Name, version, crFileName)
					} else {
						exampleCRs[version] = strings.TrimSpace(string(exampleCR))
					}
				}

				templatePath := path.Dir(configFilePath) + "/" + configuration.TemplatePath

				_, err = output.WritePage(
					crds[i],
					annotations,
					meta,
					exampleCRs,
					outputFolderPath,
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

func mustParse(fset *token.FileSet, filename string) *ast.File {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}
