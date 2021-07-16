package metadata

import (
	"io/ioutil"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v2"
)

type Root struct {
	CRDs map[string]CRDItem
}

type CRDItem struct {
	Owner     []string `yaml:"owner,omitempty"`
	Topics    []string `yaml:"topics,omitempty"`
	Providers []string `yaml:"provider,omitempty"`
	Hidden    bool     `yaml:"hidden,omitempty"`
}

func Read(path string) (*Root, error) {
	r := Root{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadFileError, err.Error())
	}

	err = yaml.UnmarshalStrict(data, &r)
	if err != nil {
		return nil, microerror.Maskf(CouldNotParseFileError, err.Error())
	}

	return &r, nil
}
