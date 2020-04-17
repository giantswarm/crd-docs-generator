package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
	blackfriday "gopkg.in/russross/blackfriday.v2"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"github.com/giantswarm/crd-docs-generator/service/git"
)

const (
	// SourceRepositoryURL is the URL to the repository defining our CRDs in Golang and YAML.
	SourceRepositoryURL = "https://github.com/giantswarm/apiextensions"

	// SourceRepositoryOrg is the Github organisation name form SourceRepositoryURL.
	SourceRepositoryOrg = "giantswarm"

	// SourceRepositoryName is the actual repo name from SourceRepositoryURL
	SourceRepositoryName = "apiextensions"
)

// CRDDocsGenerator represents an instance of this command line tool, it carries
// the cobra command which runs the process along with configuration parameters
// which come in as flags on the command line.
type CRDDocsGenerator struct {
	// Internals.
	rootCommand *cobra.Command

	// Settings/Preferences
	apiExtensionsTag string // The tag to use from the apiextensions repo when building crd documentation.
}

const (
	// Target path for our clone of the apiextensions repo.
	repoFolder = "/tmp/gitclone"

	crdFolder = repoFolder + "/config/crd/bases"

	crFolder = repoFolder + "/docs/cr"

	// Path for Markdown output.
	outputFolderPath = "./output"

	// Path for templates
	templateFolderPath = "./templates"

	// Single CRD page template filename (without path)
	outputTemplate = "crd.template"
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

// OutputData is all the data we pass to the HTML template for the CRD detail page.
type OutputData struct {
	Date                string
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
	VersionSchemas map[string]OutputSchemaVersion
}

// OutputSchemaVersion is the schema information for a specific CRD version
// we want to expose to our template.
type OutputSchemaVersion struct {
	Version    string
	Properties []SchemaProperty
	// YAML string showing an example CR.
	ExampleCR string
}

// ReadCRD reads a CRD YAML file and returns the Custom Resource Definition object it represents.
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

	// Sort properties by path.
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Path < properties[j].Path
	})

	return properties
}

func toMarkdown(input string) template.HTML {
	inputBytes := []byte(fmt.Sprintf("%s", input))
	return template.HTML(blackfriday.Run(inputBytes))
}

// WriteCRDDocs creates a CRD schema documetantation Markdown page.
func WriteCRDDocs(crd *apiextensionsv1beta1.CustomResourceDefinition, outputFolder string, repoRef string) error {
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
	data := OutputData{
		// Current date as page creation date for the front matter
		Date:                time.Now().Format("2006-01-02"),
		Group:               crd.Spec.Group,
		NamePlural:          crd.Spec.Names.Plural,
		NameSingular:        crd.Spec.Names.Singular,
		Scope:               string(crd.Spec.Scope),
		SourceRepository:    SourceRepositoryURL,
		SourceRepositoryRef: repoRef,
		Title:               crd.Spec.Names.Kind,
		Weight:              100,
		VersionSchemas:      make(map[string]OutputSchemaVersion),
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
		data.Description = crd.Spec.Validation.OpenAPIV3Schema.Description

		// Create flat attribute list from hierarchy.
		var properties []SchemaProperty
		properties = flattenProperties(crd.Spec.Validation.OpenAPIV3Schema, properties, 0, "")

		if crd.Spec.Version != "" {
			data.Versions = []string{crd.Spec.Version}
			data.VersionSchemas[crd.Spec.Version] = OutputSchemaVersion{
				Version:    crd.Spec.Version,
				Properties: properties,
			}
		} else if len(crd.Spec.Versions) == 1 {
			// There is a versions array with exactly one element, so we
			// document that and only that.
			data.Versions = []string{crd.Spec.Versions[0].Name}
			data.VersionSchemas[crd.Spec.Versions[0].Name] = OutputSchemaVersion{
				Version:    crd.Spec.Versions[0].Name,
				Properties: properties,
			}
		} else {
			fmt.Printf("WARNING: %s.%s does not have a .spec.version or .spec.versions has more than 1 element. Can't produce the expected output.\n", crd.Spec.Names.Plural, crd.Spec.Group)
		}

	} else if len(crd.Spec.Versions) > 0 {
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

			data.VersionSchemas[version.Name] = OutputSchemaVersion{
				Version:    version.Name,
				Properties: properties,
			}

			data.Versions = append(data.Versions, version.Name)
		}
	} else {
		fmt.Printf("WARNING: %s.%s does not have an OpenAPIv3 validation schema. Can't produce the expected output.\n", crd.Spec.Names.Plural, crd.Spec.Group)
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
		//return microerror.Mask(err)

		fmt.Printf("%s: %s\n", outputFile, err)
	}

	return nil
}

func main() {
	var err error

	var crdDocsGenerator CRDDocsGenerator
	{
		c := &cobra.Command{
			Use:          "crd-docs-generator",
			Short:        "crd-docs-generator is a command line tool for generating markdown files that document Giant Swarm's custom resources",
			SilenceUsage: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				return generateCrdDocs(crdDocsGenerator.apiExtensionsTag)
			},
		}

		c.PersistentFlags().StringVar(&crdDocsGenerator.apiExtensionsTag, "apiextensions-commit-ref", "master", "The git commit reference (tag, branch, commit SHA) to use from the giantswarm/apiextensions repository")
		crdDocsGenerator.rootCommand = c
	}

	if err = crdDocsGenerator.rootCommand.Execute(); err != nil {
		printStackTrace(err)
		os.Exit(1)
	}
}

func generateCrdDocs(apiExtensionsTag string) error {
	crdFiles := []string{}

	err := git.CloneRepositoryShallow(SourceRepositoryOrg, SourceRepositoryName, apiExtensionsTag, repoFolder)
	if err != nil {
		return microerror.Mask(err)
	}

	defer os.RemoveAll(repoFolder)

	err = filepath.Walk(crdFolder, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yaml") {
			fmt.Printf("Collecting file %s\n", path)
			crdFiles = append(crdFiles, path)
		}
		return nil
	})
	if err != nil {
		return microerror.Mask(err)
	}

	for _, crdFile := range crdFiles {
		fmt.Printf("Reading file %s\n", crdFile)

		crd, err := ReadCRD(crdFile)
		if err != nil {
			fmt.Printf("Something went wrong in ReadCRD: %#v\n", err)
		}

		fmt.Printf("Writing output for CRD %s\n", crd.Name)

		err = WriteCRDDocs(crd, outputFolderPath, apiExtensionsTag)
		if err != nil {
			fmt.Printf("Something went wrong in WriteCRDDocs: %#v\n", err)
		}
	}

	return nil
}

func printStackTrace(err error) {
	fmt.Println("\n--- Stack Trace ---")
	var stackedError microerror.JSONError
	json.Unmarshal([]byte(microerror.JSON(err)), &stackedError)

	for i, j := 0, len(stackedError.Stack)-1; i < j; i, j = i+1, j-1 {
		stackedError.Stack[i], stackedError.Stack[j] = stackedError.Stack[j], stackedError.Stack[i]
	}

	for _, entry := range stackedError.Stack {
		fmt.Printf("%s:%d\n", entry.File, entry.Line)
	}
}
