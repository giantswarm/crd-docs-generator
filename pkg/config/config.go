package config

import (
	"io/ioutil"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v2"
)

// FromFile represent a config file content.
type FromFile struct {
	SourceRepository *FromFileSourceRepository `yaml:"source_repository"`
}

// FromFileSourceRepository has details about the
// source repository to use for CRDs.
type FromFileSourceRepository struct {
	URL          string `yaml:"url"`
	Organization string `yaml:"organization"`
	ShortName    string `yaml:"short_name"`
	MetadataPath string `yaml:"metadata_path"`
}

// Read reads a config file and returns a struct.
func Read(path string) (*FromFile, error) {
	f := FromFile{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadConfigFileError, err.Error())
	}

	err = yaml.UnmarshalStrict(data, &f)
	if err != nil {
		return nil, microerror.Maskf(CouldNotParseConfigFileError, err.Error())
	}

	return &f, nil
}
