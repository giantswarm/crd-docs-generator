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
		return nil, microerror.Maskf(CouldNotReadCRDFileError, "%s", err.Error())
	}

	// Split by "---"
	parts := strings.Split(string(yamlBytes), "\n---\n")
	crds := []any{}

	for _, crdYAMLString := range parts {
		crdYAMLBytes := []byte(crdYAMLString)
		crd := apiextensionsv1.CustomResourceDefinition{}

		err = yaml.Unmarshal(crdYAMLBytes, &crd)
		if err != nil {
			return nil, microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
		}

		// Both Crossplane and Kratix store their CRDs in different ways
		//
		// - Crossplane is an extended CRD format
		// - Kratix is a wrapped CRD format
		if crd.GroupVersionKind().Group == "apiextensions.crossplane.io" && crd.Kind == "CompositeResourceDefinition" {
			xrd := crossplanev1.CompositeResourceDefinition{}
			err = yaml.Unmarshal(crdYAMLBytes, &xrd)
			if err != nil {
				return nil, microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
			}
			crds = append(crds, xrd)
			continue
		} else if crd.GroupVersionKind().Group == "platform.kratix.io" && crd.Kind == "Promise" {
			var iface any
			err = yaml.Unmarshal(crdYAMLBytes, &iface)
			if err != nil {
				return nil, microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
			}

			err = ParsePromise(iface, &crd)
			if err != nil {
				return nil, microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
			}
		}

		// If we had empty parts parsed, let's skip them.
		if crd.Name == "" {
			continue
		}
		crds = append(crds, crd)
	}

	return crds, nil
}

// Parse a kratix promise.
func ParsePromise(iface any, crd *apiextensionsv1.CustomResourceDefinition) (err error) {
	var (
		api          map[string]any
		spec         map[string]any
		ok           bool
		crdYAMLBytes []byte
	)

	spec, ok = iface.(map[string]any)["spec"].(map[string]any)
	if !ok {
		return microerror.Maskf(CouldNotParseCRDFileError, "%s", "kratix promise is missing its spec")
	}

	api, ok = spec["api"].(map[string]any)
	if !ok {
		return microerror.Maskf(CouldNotParseCRDFileError, "%s", "kratix promise is missing its spec.api")
	}

	crdYAMLBytes, err = yaml.Marshal(api)
	if err != nil {
		return microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
	}

	err = yaml.Unmarshal(crdYAMLBytes, &crd)
	if err != nil {
		return microerror.Maskf(CouldNotParseCRDFileError, "%s", err.Error())
	}

	return
}
