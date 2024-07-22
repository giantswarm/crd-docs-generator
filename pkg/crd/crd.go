package crd

import (
	"os"
	"strings"

	crossplanev1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Read reads a CRD YAML file and returns the Custom Resource Definition objects it represents.
func Read(filePath string) ([]any, error) {
	yamlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadCRDFileError, err.Error())
	}

	// Split by "---"
	parts := strings.Split(string(yamlBytes), "\n---\n")
	crds := []any{}

	for _, crdYAMLString := range parts {
		crdYAMLBytes := []byte(crdYAMLString)
		crd := apiextensionsv1.CustomResourceDefinition{}

		err = yaml.Unmarshal(crdYAMLBytes, &crd)
		if err != nil {
			return nil, microerror.Maskf(CouldNotParseCRDFileError, err.Error())
		}

		// If we had empty parts parsed, let's skip them.
		if crd.Name == "" {
			continue
		}

		if crd.APIVersion == "apiextensions.crossplane.io/v1" && crd.Kind == "CompositeResourceDefinition" {
			xrd := crossplanev1.CompositeResourceDefinition{}
			err = yaml.Unmarshal(crdYAMLBytes, &xrd)
			if err != nil {
				return nil, microerror.Maskf(CouldNotParseCRDFileError, err.Error())
			}
			crds = append(crds, xrd)
		} else {
			crds = append(crds, crd)
		}
	}

	return crds, nil
}
