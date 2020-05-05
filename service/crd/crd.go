package crd

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/giantswarm/microerror"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Read reads a CRD YAML file and returns the Custom Resource Definition object it represents.
func Read(inputFile string) (*apiextensionsv1.CustomResourceDefinition, error) {
	crd := &apiextensionsv1.CustomResourceDefinition{}

	yamlBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadCRDFileError, err.Error())
	}

	err = yaml.Unmarshal(yamlBytes, crd)
	if err != nil {
		return nil, microerror.Maskf(CouldNotParseCRDFileError, err.Error())
	}

	return crd, nil
}
