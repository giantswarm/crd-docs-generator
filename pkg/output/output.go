package output

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/giantswarm/microerror"
	blackfriday "github.com/russross/blackfriday/v2"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/crd-docs-generator/pkg/jsonschema"
	"github.com/giantswarm/crd-docs-generator/pkg/metadata"
)

// CRDAnnotationSupport represents the release and
type CRDAnnotationSupport struct {
	Annotation    string
	APIVersion    string
	CRDName       string
	Release       string
	Documentation string
}

// PageData is all the data we pass to the HTML template for the CRD detail page.
type PageData struct {
	Description         string
	Metadata            metadata.CRDItem
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
	Annotations []CRDAnnotationSupport
}

// WritePage creates a CRD schema documentation Markdown page.
func WritePage(crd apiextensionsv1.CustomResourceDefinition,
	annotations []CRDAnnotationSupport,
	md metadata.CRDItem,
	crFolder,
	outputFolder,
	repoURL,
	repoRef,
	templatePath string) (string, error) {

	templateCode, err := ioutil.ReadFile(templatePath)
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

	// Case B: CRD contains multiple versions and schemas.
	for _, version := range crd.Spec.Versions {
		if !version.Served && !version.Storage {
			// Neither stored nore served means that this version
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
			Annotations: sortAnnotations(filterAnnotations(annotations, crd.GetObjectMeta().GetName(), version.Name)),
		}

		data.Versions = append(data.Versions, version.Name)
	}

	// Try to read example CRs for all versions.
	for _, version := range data.Versions {
		crFileName := fmt.Sprintf("%s/%s_%s_%s.yaml", crFolder, crd.Spec.Group, version, crd.Spec.Names.Singular)
		exampleCR, err := ioutil.ReadFile(crFileName)
		if err != nil {
			fmt.Printf("%s - CR example is missing\n", crd.Name)
		} else {
			outputSchema := data.VersionSchemas[version]
			outputSchema.ExampleCR = string(exampleCR)
			data.VersionSchemas[version] = outputSchema
		}
	}

	// Name output file after full CRD name.
	outputFile := outputFolder + "/" + crd.Spec.Names.Plural + "." + crd.Spec.Group + ".md"

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

func filterAnnotations(annotations []CRDAnnotationSupport, CRDName string, APIVersion string) []CRDAnnotationSupport {
	var result []CRDAnnotationSupport

	for _, annotation := range annotations {
		if annotation.CRDName == CRDName && annotation.APIVersion == APIVersion {
			result = append(result, annotation)
		}
	}

	return result
}

func sortAnnotations(annotations []CRDAnnotationSupport) []CRDAnnotationSupport {
	sort.Slice(annotations, func(i, j int) bool {
		return annotations[i].Annotation < annotations[j].Annotation
	})

	return annotations
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
