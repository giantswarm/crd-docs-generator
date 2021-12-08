package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
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

	annotationsFolder = repoFolder + "/pkg/annotation"

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

	crdFiles := []string{}
	annotationFiles := []string{}

	// TODO: Loop over configured repositories
	defer os.RemoveAll(repoFolder)
	for _, sourceRepo := range configuration.SourceRepositories {
		clonePath := repoFolder + "/" + sourceRepo.Organization + "/" + sourceRepo.ShortName
		// Clone the repositories containing CRDs
		err = git.CloneRepositoryShallow(
			sourceRepo.Organization,
			sourceRepo.ShortName,
			sourceRepo.CommitReference,
			clonePath)
		if err != nil {
			return microerror.Mask(err)
		}

		// Collect annotation info
		err = filepath.Walk(annotationsFolder, func(path string, info os.FileInfo, err error) error {
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
				crdFiles = append(crdFiles, path)
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
				crdFiles = append(crdFiles, path)
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
					fmt.Printf("%s - %s - annotation YAML docs missing\n", annotationFile, constant.Names[0])
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

		for _, crdFile := range crdFiles {
			crds, err := crd.Read(crdFile)
			if err != nil {
				fmt.Printf("Something went wrong in crd.Read: %#v\n", err)
			}

			for i := range crds {
				_, exists := crdNames[crds[i].Name]
				if exists {
					log.Printf("WARN - repo %s - provides CRD %s which is already added - skipping", sourceRepo.ShortName, crds[i].Name)
					continue
				}
				crdNames[crds[i].Name] = true
				// Skip hidden CRDs and CRDs with missing metadata
				meta, ok := sourceRepo.Metadata[crds[i].Name]
				if !ok {
					fmt.Printf("%s - metadata is missing, skipping\n", crds[i].Name)
					continue
				}
				if meta.Hidden {
					fmt.Printf("%s - is hidden explicitly, skipping\n", crds[i].Name)
					continue
				}

				templatePath := path.Dir(configFilePath) + "/" + configuration.TemplatePath

				_, err = output.WritePage(
					crds[i],
					annotations,
					meta,
					crFolder,
					outputFolderPath,
					sourceRepo.URL,
					sourceRepo.CommitReference,
					templatePath)
				if err != nil {
					fmt.Printf("Something went wrong in WriteCRDDocs: %#v\n", err)
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
