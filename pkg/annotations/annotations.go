package annotations

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v2"
)

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
			annotation := Annotation{}

			err := yaml.UnmarshalStrict([]byte(constant.Doc), &annotation)
			if err != nil {
				log.Printf("WARN - Annotation in %s named %q does not provide compatible YAML docs", annotationFile, constant.Names[0])
			}

			if annotation.Documentation != "" {
				for _, crdAPI := range annotation.Support {
					rawAnnotation := constant.Decl.Specs[0].(*ast.ValueSpec).Values[0].(*ast.BasicLit).Value
					annotationValue := strings.Replace(rawAnnotation, "\"", "", -1)

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
