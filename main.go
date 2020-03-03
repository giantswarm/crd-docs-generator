package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
	blackfriday "gopkg.in/russross/blackfriday.v2"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

const (
	// Path for the CRD YAML files folder.
	inputFolderPath = "./crd"

	// Path for Markdown output.
	outputFolderPath = "./output"

	// Path for templates
	templateFolderPath = "./templates"

	// Single CRD page template filename (without path)
	outputTemplate = "crd.template"
)

var (
	config = []ConfigCRDItem{
		ConfigCRDItem{
			FileNamePrefix: "app",
		},
		ConfigCRDItem{
			FileNamePrefix: "appcatalog",
		},
		ConfigCRDItem{
			FileNamePrefix: "awscluster",
		},
		ConfigCRDItem{
			FileNamePrefix: "awscontrolplane",
		},
		ConfigCRDItem{
			FileNamePrefix: "awsmachinedeployment",
		},
		ConfigCRDItem{
			FileNamePrefix: "chart",
		},
		ConfigCRDItem{
			FileNamePrefix: "g8scontrolplane",
		},
		ConfigCRDItem{
			FileNamePrefix: "release",
		},
		ConfigCRDItem{
			FileNamePrefix: "releasecycle",
		},
	}
)

// ConfigCRDItem configues one specific CRD to document.
type ConfigCRDItem struct {
	// First part of the CRD file name to read as input and
	// the Markdown file to crete as output.
	FileNamePrefix string
}

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

type OutputData struct {
	Date         string
	Description  string
	Group        string
	NamePlural   string
	NameSingular string
	Scope        string
	Title        string
	Weight       int
	// Version names.
	Versions []string
	// Schema per version.
	VersionSchemas map[string]OutputSchemaVersion
}

type OutputSchemaVersion struct {
	Version    string
	Properties []SchemaProperty
}

// ReadCRD parses a CRD YAML file and creates markdown documentation.
func ReadCRD(inputFile string) (*apiextensionsv1beta1.CustomResourceDefinition, error) {
	crd := &apiextensionsv1beta1.CustomResourceDefinition{}

	yamlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	err = yaml.Unmarshal(yamlBytes, crd)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return crd, nil
}

// flattenProperties recurses over all properties of a JSON Schema
// and returns a flat slice of the elements we need for our output.
func flattenProperties(schema *apiextensionsv1beta1.JSONSchemaProps, properties []SchemaProperty, depth int8, pathPrefix string) []SchemaProperty {
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

	return properties
}

func toMarkdown(input string) template.HTML {
	inputBytes := []byte(fmt.Sprintf("%s", input))
	return template.HTML(blackfriday.Run(inputBytes))
}

// WriteCRDDocs creates a CRD schema documetantation Markdown page.
func WriteCRDDocs(crd *apiextensionsv1beta1.CustomResourceDefinition, outputFile string) error {
	templateCode, err := ioutil.ReadFile(templateFolderPath + "/" + outputTemplate)
	if err != nil {
		return microerror.Mask(err)
	}

	// Add custom functions support for our template.
	funcMap := template.FuncMap{
		// Treat given test as Markdown and convert to HTML.
		"markdown": toMarkdown,
		// Join strings by separator
		"join": strings.Join,
	}

	// Read our output template.
	tpl := template.Must(template.New("schemapage").Funcs(funcMap).Parse(string(templateCode)))

	// Collect values to pass to our output template.
	data := OutputData{
		// Current date as page creation date for the front matter
		Date:           time.Now().Format("2006-01-02"),
		Group:          crd.Spec.Group,
		NamePlural:     crd.Spec.Names.Plural,
		NameSingular:   crd.Spec.Names.Singular,
		Scope:          string(crd.Spec.Scope),
		Title:          crd.Spec.Names.Kind,
		Weight:         100,
		VersionSchemas: make(map[string]OutputSchemaVersion),
	}

	// We handle two very different cases here and bring them to a unififed output structure.
	//
	// A: CRD contains only one version defined as .spec.version and .spec.validation contains
	// the schema.
	//
	// B: CRD contains multiple schemas under .spec.versions[*] and schema under
	// .spec.versions[*].schema
	//
	if crd.Spec.Validation != nil {
		// Case A: CRD contains only one version defined as .spec.version.

		// Create flat attribute list from hierarchy.
		var properties []SchemaProperty
		properties = flattenProperties(crd.Spec.Validation.OpenAPIV3Schema, properties, 0, "")

		// Sort properties by path.
		sort.Slice(properties, func(i, j int) bool {
			return properties[i].Path < properties[j].Path
		})

		if crd.Spec.Version != "" {
			data.Versions = []string{crd.Spec.Version}
			data.VersionSchemas[crd.Spec.Version] = OutputSchemaVersion{
				Version:    crd.Spec.Version,
				Properties: properties,
			}
		} else {
			fmt.Printf("WARNING: %s.%s does not have a .spec.version. Can't produce the expected output.\n", crd.Spec.Names.Plural, crd.Spec.Group)
		}

	} else if len(crd.Spec.Versions) > 0 {
		// Case B: CRD contains multiple versions and schemas.

		for _, version := range crd.Spec.Versions {
			var properties []SchemaProperty
			properties = flattenProperties(version.Schema.OpenAPIV3Schema, properties, 0, "")

			data.VersionSchemas[version.Name] = OutputSchemaVersion{
				Version:    version.Name,
				Properties: properties,
			}

			data.Versions = append(data.Versions, version.Name)
		}
	} else {
		fmt.Printf("WARNING: %s.%s does not have an OpenAPIv3 validation schema. Can't produce the expected output.\n", crd.Spec.Names.Plural, crd.Spec.Group)
	}

	handler, err := os.Create(outputFile)
	if err != nil {
		return microerror.Mask(err)
	}

	err = tpl.Execute(handler, data)
	if err != nil {

		// TODO: return error
		//return microerror.Mask(err)

		fmt.Printf("%s: %s\n", outputFile, err)
	}

	return nil
}

func main() {
	for _, config := range config {
		inputFile := inputFolderPath + "/" + config.FileNamePrefix + ".yaml"

		crd, err := ReadCRD(inputFile)
		if err != nil {
			fmt.Printf("Something went wrong in ReadCRD: %#v", err)
		}

		outputFile := outputFolderPath + "/" + config.FileNamePrefix + ".md"

		err = WriteCRDDocs(crd, outputFile)
		if err != nil {
			fmt.Printf("Something went wrong in WriteCRDDocs: %#v", err)
		}
	}

}
