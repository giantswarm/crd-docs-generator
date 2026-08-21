package jsonschema

import (
	"reflect"
	"testing"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// JSON Schema type and property names repeated across the fixtures below.
const (
	typeObject = "object"
	typeString = "string"

	requiredStringProp = "required_string"
)

func TestFlatten(t *testing.T) {
	type args struct {
		schema     apiextensionsv1.JSONSchemaProps
		properties []Property
		depth      int8
		pathPrefix string
	}
	tests := []struct {
		name string
		args args
		want []Property
	}{
		{
			name: "Nested schema",
			args: args{
				schema: apiextensionsv1.JSONSchemaProps{
					ID:          "root",
					Description: "top description",
					Type:        typeObject,
					Required:    []string{requiredStringProp},
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						requiredStringProp: {
							ID:          "the_id",
							Description: "A required string property",
							Type:        typeString,
						},
						"optional_array": {
							Type:        "array",
							Description: "An optional array property",
							Items: &apiextensionsv1.JSONSchemaPropsOrArray{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Description: "Array item",
									Type:        typeString,
								},
								JSONSchemas: []apiextensionsv1.JSONSchemaProps{},
							},
						},
						"optional_object": {
							Type: typeObject,
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"nested_number": {
									Type: "number",
								},
							},
						},
					},
				},
				properties: []Property{},
				depth:      0,
				pathPrefix: "",
			},
			want: []Property{
				{
					Depth:       0,
					Path:        ".optional_array",
					Name:        "optional_array",
					Type:        "array",
					Description: "An optional array property",
					Required:    false,
				},
				{
					Depth:       1,
					Path:        ".optional_array[*]",
					Name:        "optional_array[*]",
					Type:        typeString,
					Description: "Array item",
					Required:    false,
				},
				{
					Depth:       0,
					Path:        ".optional_object",
					Name:        "optional_object",
					Type:        typeObject,
					Description: "",
					Required:    false,
				},
				{
					Depth:       1,
					Path:        ".optional_object.nested_number",
					Name:        "nested_number",
					Type:        "number",
					Description: "",
					Required:    false,
				},
				{
					Depth:       0,
					Path:        ".required_string",
					Name:        requiredStringProp,
					Type:        typeString,
					Description: "A required string property",
					Required:    true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.schema, tt.args.properties, tt.args.depth, tt.args.pathPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
