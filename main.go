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
	"path/filepath"

	"strings"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

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

	// Path to the CRD page template file
	templateFilePath string

	// git reference (tag, commit SHA, branch name) to check out for the source repository
	sourceCommitRef string
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

	crdFolder = repoFolder + "/config/crd"

	crFolder = repoFolder + "/docs/cr"

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
				return generateCrdDocs(crdDocsGenerator.configFilePath,
					crdDocsGenerator.templateFilePath,
					crdDocsGenerator.sourceCommitRef)
			},
		}

		c.PersistentFlags().StringVar(&crdDocsGenerator.configFilePath, "config", "./config.yaml", "Path to the configuration file.")
		c.PersistentFlags().StringVar(&crdDocsGenerator.templateFilePath, "template", "./templates/crd.template", "Path to the CRD page template file.")
		c.PersistentFlags().StringVar(&crdDocsGenerator.sourceCommitRef, "commit-reference", "main", "Commit SHA, tag or branch name to use of the CRD source repository.")
		crdDocsGenerator.rootCommand = c
	}

	if err = crdDocsGenerator.rootCommand.Execute(); err != nil {
		printStackTrace(err)
		os.Exit(1)
	}
}

// generateCrdDocs is the function called from our main CLI command.
func generateCrdDocs(configFilePath, templatePath, commitRef string) error {
	configuration, err := config.Read(configFilePath)
	if err != nil {
		return microerror.Mask(err)
	}

	crdFiles := []string{}
	annotationFiles := []string{}

	err = git.CloneRepositoryShallow(
		configuration.SourceRepository.Organization,
		configuration.SourceRepository.ShortName,
		commitRef,
		repoFolder)
	if err != nil {
		return microerror.Mask(err)
	}

	defer os.RemoveAll(repoFolder)

	err = filepath.Walk(annotationsFolder, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		if strings.HasSuffix(path, ".go") {
			fmt.Printf("Collecting file %s\n", path)
			annotationFiles = append(annotationFiles, path)
		}
		return nil
	})
	if err != nil {
		return microerror.Mask(err)
	}

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
				log.Printf("%s has no yaml documentation", constant.Names[0])
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
				annotations,
				crFolder,
				outputFolderPath,
				configuration.SourceRepository.URL,
				commitRef,
				templatePath)
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
