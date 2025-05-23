package output

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/giantswarm/microerror"
	blackfriday "github.com/russross/blackfriday/v2"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/crd-docs-generator/pkg/annotations"
	"github.com/giantswarm/crd-docs-generator/pkg/config"
	"github.com/giantswarm/crd-docs-generator/pkg/jsonschema"
)

// PageData is all the data we pass to the HTML template for the CRD detail page.
type PageData struct {
	Description         string
	Metadata            config.CRDItem
	Group               string
	NamePlural          string
	NameSingular        string
	Scope               string
	SourceRepository    string
	SourceRepositoryRef string
	Title               string
	Weight              int
	// Version names.
	Versions []string
	// Schema per version.
	VersionSchemas map[string]SchemaVersion
}

// SchemaVersion is the schema information for a specific CRD version
// we want to expose to our template.
type SchemaVersion struct {
	Version    string
	Properties []jsonschema.Property
	// YAML string showing an example CR.
	ExampleCR   string
	Annotations []annotations.CRDAnnotationSupport
}

// WritePage creates a CRD schema documentation Markdown page.
func WritePage(crd apiextensionsv1.CustomResourceDefinition,
	crdAnnotations []annotations.CRDAnnotationSupport,
	md config.CRDItem,
	examplesCRs map[string]string,
	outputFolder,
	repoURL,
	repoRef,
	templatePath string) (string, error) {

	templatePath = filepath.Clean(templatePath)
	templateCode, err := os.ReadFile(templatePath)
	if err != nil {
		return "", microerror.Maskf(cannotOpenTemplate, "Could not read template file %s: %s", templatePath, err)
	}

	// Add custom functions support for our template.
	funcMap := sprig.FuncMap()
	// Treat given test as Markdown and convert to HTML.
	funcMap["markdown"] = toMarkdown
	// Join strings by separator
	funcMap["join"] = strings.Join
	// Return raw string
	funcMap["raw"] = rawString

	// Read our output template.
	tpl := template.Must(template.New("schemapage").Funcs(funcMap).Parse(string(templateCode)))

	// Collect values to pass to our output template.
	data := PageData{
		Group:               crd.Spec.Group,
		Metadata:            md,
		NamePlural:          crd.Spec.Names.Plural,
		NameSingular:        crd.Spec.Names.Singular,
		Scope:               string(crd.Spec.Scope),
		SourceRepository:    repoURL,
		SourceRepositoryRef: repoRef,
		Title:               crd.Spec.Names.Kind,
		Weight:              100,
		VersionSchemas:      make(map[string]SchemaVersion),
	}

	// Iterate schema versions
	for _, version := range crd.Spec.Versions {
		if !version.Served && !version.Storage {
			// Neither stored nor served means that this version
			// can be skipped.
			continue
		}

		// Get the first non-empty top level description and use it as the
		// CRD description.
		if data.Description == "" && version.Schema != nil {
			data.Description = version.Schema.OpenAPIV3Schema.Description
		}

		var properties []jsonschema.Property

		if version.Schema != nil && version.Schema.OpenAPIV3Schema != nil {
			properties = jsonschema.Flatten(*version.Schema.OpenAPIV3Schema, properties, 0, "")
		}

		data.VersionSchemas[version.Name] = SchemaVersion{
			Version:     version.Name,
			Properties:  properties,
			Annotations: annotations.FilterForCRD(crdAnnotations, crd.Name, version.Name),
		}

		data.Versions = append(data.Versions, version.Name)
	}

	sort.Strings(data.Versions)

	// Add example CRs
	for _, version := range data.Versions {
		exampleCR, ok := examplesCRs[version]
		if ok {
			outputSchema := data.VersionSchemas[version]
			outputSchema.ExampleCR = exampleCR + "\n"
			data.VersionSchemas[version] = outputSchema
		}
	}

	// Name output file after full CRD name.
	outputFile := outputFolder + "/" + crd.Spec.Names.Plural + "." + crd.Spec.Group + ".md"

	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.MkdirAll(outputFolder, 0750)
		if err != nil {
			return "", microerror.Mask(err)
		}
	}

	outputFile = filepath.Clean(outputFile)
	handler, err := os.Create(outputFile)
	if err != nil {
		return "", microerror.Mask(err)
	}

	err = tpl.Execute(handler, data)
	if err != nil {

		// TODO: return error
		// return microerror.Mask(err)

		fmt.Printf("%s: %s\n", outputFile, err)
	}

	return outputFile, nil
}

func toMarkdown(input string) template.HTML {
	inputBytes := []byte(input)
	// To mitigate gosec "this method will not auto-escape HTML. Verify data is well formed"
	// #nosec G203
	return template.HTML(blackfriday.Run(inputBytes))
}

func rawString(input string) template.HTML {
	// To mitigate gosec "this method will not auto-escape HTML. Verify data is well formed"
	// #nosec G203
	return template.HTML(input)
}
