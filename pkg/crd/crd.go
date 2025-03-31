package crd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Read reads a CRD YAML file and returns the Custom Resource Definition objects it represents.
func Read(filePath string) ([]apiextensionsv1.CustomResourceDefinition, error) {
	filePath = filepath.Clean(filePath)
	yamlBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadCRDFileError, "%s", err.Error())
	}

	// Split by "---"
	parts := strings.Split(string(yamlBytes), "\n---\n")
	crds := []apiextensionsv1.CustomResourceDefinition{}

	for _, crdYAMLString := range parts {
		crdYAMLBytes := []byte(crdYAMLString)
		crd := apiextensionsv1.CustomResourceDefinition{}

		err = yaml.Unmarshal(crdYAMLBytes, &crd)
		if err != nil {
			return nil, microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
		}

		// If we had empty parts parsed, let's skip them.
		if crd.Name == "" {
			continue
		}

		crds = append(crds, crd)
	}

	return crds, nil
}
