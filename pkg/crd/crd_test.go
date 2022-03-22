package crd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestRead(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []apiextensionsv1.CustomResourceDefinition
		wantErr bool
	}{
		{
			name: "file with one CRD with one version and some schema",
			args: args{
				filePath: "testdata/awsclusterconfig.yaml",
			},
			want: []apiextensionsv1.CustomResourceDefinition{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "CustomResourceDefinition",
						APIVersion: "apiextensions.k8s.io/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name: "awsclusterconfigs.core.giantswarm.io",
						Annotations: map[string]string{
							"controller-gen.kubebuilder.io/version": "v0.2.4",
						},
					},
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Group: "core.giantswarm.io",
						Names: apiextensionsv1.CustomResourceDefinitionNames{
							Plural:     "awsclusterconfigs",
							Singular:   "awsclusterconfig",
							Kind:       "AWSClusterConfig",
							ListKind:   "AWSClusterConfigList",
							Categories: []string{"aws", "giantswarm"},
						},
						Scope: "Namespaced",
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
							{
								Name:    "v1alpha1",
								Served:  true,
								Storage: true,
								Schema: &apiextensionsv1.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
										Description: "AWSClusterConfig used to represent workload cluster configuration in earlier releases. Deprecated.",
										Type:        "object",
										Required: []string{
											"spec",
										},
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"apiVersion": {
												Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
												Type:        "string",
											},
											"kind": {
												Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Type:        "string",
											},
											"spec": {
												Type:     "object",
												Required: []string{"someString"},
												Properties: map[string]apiextensionsv1.JSONSchemaProps{
													"someString": {
														Type: "string",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Read() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
