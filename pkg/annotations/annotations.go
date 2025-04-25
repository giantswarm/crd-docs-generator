package annotations

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/goccy/go-yaml"
)

const CRD_DOCS_GENERATOR = "CRD_DOCS_GENERATOR"

type Annotation struct {
	Documentation string
	Support       []AnnotationSupportRelease
}

type AnnotationSupportRelease struct {
	Release    string
	APIVersion string
	CRD        string
}

// CRDAnnotationSupport is a flattened combination of
// annotation details and the CRD they are applicable to.
type CRDAnnotationSupport struct {
	Annotation    string
	CRDName       string
	CRDVersion    string
	Release       string
	Documentation string
}

// Collect finds all annotations in a folder and
// returns them.
func Collect(startPath string) ([]CRDAnnotationSupport, error) {
	var annotations []CRDAnnotationSupport

	files, err := findFiles(startPath)
	if err != nil {
		return annotations, microerror.Mask(err)
	}

	for _, annotationFile := range files {
		fset := token.NewFileSet()
		files := []*ast.File{
			mustParse(fset, annotationFile),
		}

		p, err := doc.NewFromFiles(fset, files, "github.com/giantswarm/extract-go-doc/p")
		if err != nil {
			panic(err)
		}

		for _, constant := range p.Consts {
			annotation, err := parseAnnotation(constant.Doc)
			if err != nil {
				log.Printf("WARN - Annotation in %s named %q does not provide compatible YAML docs", annotationFile, constant.Names[0])
				continue
			}

			if annotation.Documentation != "" {
				for _, crdAPI := range annotation.Support {
					rawAnnotation := constant.Decl.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value
					annotationValue := strings.ReplaceAll(rawAnnotation, "\"", "")

					annotations = append(annotations, CRDAnnotationSupport{
						Annotation:    annotationValue,
						CRDName:       crdAPI.CRD,
						CRDVersion:    crdAPI.APIVersion,
						Release:       crdAPI.Release,
						Documentation: annotation.Documentation,
					})
				}
			}
		}
	}

	return Sort(annotations), nil
}

func parseAnnotation(rawAnnotation string) (*Annotation, error) {
	lines := strings.Split(rawAnnotation, "\n")

	crdCocsLineIndex := getCrdDocsLineIndex(lines)
	if crdCocsLineIndex == -1 {
		return nil, fmt.Errorf("no %s line found", CRD_DOCS_GENERATOR)
	}

	lines = lines[crdCocsLineIndex+1:]
	lines = unIndent(lines)
	rawAnnotation = strings.Join(lines, "\n")

	annotation := &Annotation{}
	reader := bytes.NewReader([]byte(rawAnnotation))
	// Fail on unknown fields.
	decoder := yaml.NewDecoder(reader, yaml.DisallowUnknownField())
	err := decoder.Decode(annotation)
	if err != nil {
		return nil, err
	}

	return annotation, nil
}

func getCrdDocsLineIndex(lines []string) int {
	for i, line := range lines {
		if strings.Contains(line, CRD_DOCS_GENERATOR) {
			return i
		}
	}

	return -1
}

// removes leading tabs from each line
func unIndent(lines []string) []string {
	var result []string
	for _, line := range lines {
		result = append(result, strings.TrimPrefix(line, "\t"))
	}

	return result
}

func findFiles(startPath string) ([]string, error) {
	annotationFiles := []string{}
	err := filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			annotationFiles = append(annotationFiles, path)
		}
		return nil
	})
	if err != nil {
		return annotationFiles, microerror.Mask(err)
	}

	return annotationFiles, nil
}

func mustParse(fset *token.FileSet, filename string) *ast.File {
	filename = filepath.Clean(filename)
	src, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}

func FilterForCRD(annotations []CRDAnnotationSupport, crdName string, version string) []CRDAnnotationSupport {
	var result []CRDAnnotationSupport

	for _, annotation := range annotations {
		if annotation.CRDName != crdName {
			continue
		}
		if version != "" && annotation.CRDVersion != version {
			continue
		}
		result = append(result, annotation)
	}

	return result
}

func Sort(annotations []CRDAnnotationSupport) []CRDAnnotationSupport {
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i].Annotation < annotations[j].Annotation
	})

	return annotations
}
