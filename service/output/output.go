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
	blackfriday "gopkg.in/russross/blackfriday.v2"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// SchemaProperty is a simplistic, flattened representation of a property
// in a JSON Schema, without the recursion and containing only the elements
// we intend to expose in our output.
type SchemaProperty struct {
	// The depth of the item in the JSONPath hierarchy
	Depth int8
	// Path is the full JSONpath path of the attribute, e. g. ".spec.version".
	Path string
	// Name is the attribute name.
	Name string
	// Type is the textual representaiton of the type ("object", "array", "number", "string", "boolean").
	Type string
	// Description is a user-friendly description of the attribute.
	Description string
	// Required specifies whether the property is required.
	Required bool
}

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
	Properties []SchemaProperty
	// YAML string showing an example CR.
	ExampleCR   string
	Annotations []CRDAnnotationSupport
}

// WritePage creates a CRD schema documentation Markdown page.
func WritePage(crd *apiextensionsv1.CustomResourceDefinition, annotations []CRDAnnotationSupport, crFolder, outputFolder, repoURL, repoRef, templateFolderPath, outputTemplate string) error {
	templateCode, err := ioutil.ReadFile(templateFolderPath + "/" + outputTemplate)
	if err != nil {
		return microerror.Mask(err)
	}

	// Add custom functions support for our template.
	funcMap := sprig.FuncMap()
	// Treat given test as Markdown and convert to HTML.
	funcMap["markdown"] = toMarkdown
	// Join strings by separator
	funcMap["join"] = strings.Join

	// Read our output template.
	tpl := template.Must(template.New("schemapage").Funcs(funcMap).Parse(string(templateCode)))

	// Collect values to pass to our output template.
	data := PageData{
		Group:               crd.Spec.Group,
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

		var properties []SchemaProperty

		if version.Schema != nil && version.Schema.OpenAPIV3Schema != nil {
			properties = flattenProperties(version.Schema.OpenAPIV3Schema, properties, 0, "")
		}

		data.VersionSchemas[version.Name] = SchemaVersion{
			Version:     version.Name,
			Properties:  properties,
			Annotations: filterAnnotations(annotations, crd.GetObjectMeta().GetName(), version.Name),
		}

		data.Versions = append(data.Versions, version.Name)
	}

	// Try to read example CRs for all versions.
	for _, version := range data.Versions {
		crFileName := fmt.Sprintf("%s/%s_%s_%s.yaml", crFolder, crd.Spec.Group, version, crd.Spec.Names.Singular)
		exampleCR, err := ioutil.ReadFile(crFileName)
		if err != nil {
			fmt.Printf("Error when reading example CR file: %s\n", err)
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
		return microerror.Mask(err)
	}

	err = tpl.Execute(handler, data)
	if err != nil {

		// TODO: return error
		// return microerror.Mask(err)

		fmt.Printf("%s: %s\n", outputFile, err)
	}

	return nil
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

func toMarkdown(input string) template.HTML {
	inputBytes := []byte(input)
	// To mitigate gosec "this method will not auto-escape HTML. Verify data is well formed"
	// #nosec G203
	return template.HTML(blackfriday.Run(inputBytes))
}

// flattenProperties recurses over all properties of a JSON Schema
// and returns a flat slice of the elements we need for our output.
func flattenProperties(schema *apiextensionsv1.JSONSchemaProps, properties []SchemaProperty, depth int8, pathPrefix string) []SchemaProperty {
	// Capture names of required properties.
	requiredProps := make(map[string]bool)
	for _, p := range schema.Required {
		requiredProps[p] = true
	}

	// Collect reduced property info.
	for propname, schemaProps := range schema.Properties {
		path := fmt.Sprintf("%s.%s", pathPrefix, propname)

		required := false
		if _, ok := requiredProps[propname]; ok {
			required = true
		}

		property := SchemaProperty{
			Depth:       depth,
			Name:        propname,
			Path:        path,
			Description: schemaProps.Description,
			Type:        schemaProps.Type,
			Required:    required,
		}

		properties = append(properties, property)

		if len(schemaProps.Properties) > 0 {
			properties = flattenProperties(&schemaProps, properties, depth+1, path)
		}

		if schemaProps.Type == "array" && schemaProps.Items != nil {
			// Add description of array member type
			property := SchemaProperty{
				Depth:       depth + 1,
				Name:        propname + "[*]",
				Path:        path + "[*]",
				Description: schemaProps.Items.Schema.Description,
				Type:        schemaProps.Items.Schema.Type,
			}
			properties = append(properties, property)

			// Collect sub items properties
			properties = flattenProperties(schemaProps.Items.Schema, properties, depth+2, path+"[*]")
		}
	}

	// Sort properties by path.
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Path < properties[j].Path
	})

	return properties
}
