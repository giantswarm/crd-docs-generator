package jsonschema

import (
	"fmt"
	"sort"
	"strings"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

const arrayItemIndicator = "[*]"

// Property is a simplistic, flattened representation of a property
// in a JSON Schema, without the recursion and containing only the elements
// we intend to expose in our output.
type Property struct {
	// The default value of the attribute.
	Default string
	// The depth of the item in the JSONPath hierarchy
	Depth int8
	// Description is a user-friendly description of the attribute.
	Description string
	// Enum is a list of possible values for the attribute.
	Enum []string
	// Immutability specifies how the attribute is immutable.
	Immutability string
	// MinLength is the minimum length of the attribute.
	MinLength int
	// MaxLength is the maximum length of the attribute.
	MaxLength int
	// Name is the attribute name.
	Name string
	// Path is the full JSONpath path of the attribute, e. g. ".spec.version".
	Path string
	// Pattern is the validation pattern required for this attribute.
	Pattern string
	// Required specifies whether the property is required.
	Required bool
	// Type is the textual representaiton of the type ("object", "array", "number", "string", "boolean").
	Type string
}

// Flatten recurses over all properties of a JSON Schema
// and returns a flat slice of the elements we need for our output.
func Flatten(schema apiextensionsv1.JSONSchemaProps, properties []Property, depth int8, pathPrefix string) []Property {
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

		var minLen int = 0
		if schemaProps.MinItems != nil {
			minLen = int(*schemaProps.MinItems)
		}

		var maxLen int = -1
		if schemaProps.MaxItems != nil {
			maxLen = int(*schemaProps.MaxItems)
		}

		if schemaProps.Type == "map" {
			minLen = int(*schemaProps.MinProperties)
			maxLen = int(*schemaProps.MaxProperties)
		}

		var enum []string
		for _, e := range schemaProps.Enum {
			var v string = string(e.Raw)
			v = strings.ReplaceAll(v, "\"", "")
			enum = append(enum, v)
		}

		pattern := strings.ReplaceAll(schemaProps.Pattern, "|", "\\|")
		description := strings.ReplaceAll(schemaProps.Description, "\n\n", "\n")

		property := Property{
			Depth:        depth,
			Description:  description,
			Enum:         enum,
			Immutability: ImmutabilityRules(schemaProps.XValidations),
			MinLength:    minLen,
			MaxLength:    maxLen,
			Name:         propname,
			Path:         path,
			Pattern:      pattern,
			Required:     required,
			Type:         schemaProps.Type,
		}

		if schemaProps.Default != nil {
			property.Default = string(schemaProps.Default.Raw)
			property.Default = strings.ReplaceAll(property.Default, "\"", "")
		}

		properties = append(properties, property)

		if len(schemaProps.Properties) > 0 {
			properties = Flatten(schemaProps, properties, depth+1, path)
		}

		if schemaProps.Type == "array" && schemaProps.Items != nil {
			var enum []string
			for _, e := range schemaProps.Enum {
				var v string = string(e.Raw)
				v = strings.ReplaceAll(v, "\"", "")
				enum = append(enum, v)
			}

			pattern := strings.ReplaceAll(schemaProps.Items.Schema.Pattern, "|", "\\|")
			description := strings.ReplaceAll(schemaProps.Items.Schema.Description, "\n\n", "\n")

			// Add description of array member type
			property := Property{
				Depth:        depth + 1,
				Description:  description,
				Enum:         enum,
				Immutability: ImmutabilityRules(schemaProps.Items.Schema.XValidations),
				Name:         propname + arrayItemIndicator,
				Path:         path + arrayItemIndicator,
				Pattern:      pattern,
				Type:         schemaProps.Items.Schema.Type,
			}

			if schemaProps.Items.Schema.Default != nil {
				property.Default = string(schemaProps.Items.Schema.Default.Raw)
				property.Default = strings.ReplaceAll(property.Default, "\"", "")
			}
			properties = append(properties, property)

			// Collect sub items properties
			properties = Flatten(*schemaProps.Items.Schema, properties, depth+2, path+arrayItemIndicator)
		}
	}

	// Sort properties by path.
	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Path < properties[j].Path
	})

	return properties
}

func ImmutabilityRules(rules apiextensionsv1.ValidationRules) string {
	var immutable string = "None"
	for _, v := range rules {
		switch v.Rule {
		case "self == oldSelf":
			immutable = "immutable"
		case "self >= oldSelf":
			immutable = "increment only"
		case "oldSelf.all(x, x in self)":
			immutable = "append only"
		case "oldSelf.all(key, key in self && self[key] == oldSelf[key])":
			immutable = "append only, immutable keys"
		}
	}
	return immutable
}
