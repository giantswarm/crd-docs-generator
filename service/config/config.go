package config

import (
	"io/ioutil"
	"sort"

	"github.com/giantswarm/microerror"
	"gopkg.in/yaml.v2"
)

// FromFile represent a config file content.
type FromFile struct {
	SourceRepository *FromFileSourceRepository `yaml:"source_repository"`
	SkipCRDs         []string                  `yaml:"skip_crds"`
}

// FromFileSourceRepository has details about the
// source repository to use for CRDs.
type FromFileSourceRepository struct {
	URL             string `yaml:"url"`
	Organization    string `yaml:"organization"`
	ShortName       string `yaml:"short_name"`
	CommitReference string `yaml:"commit_reference"`
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

	// Stable sorting for reproducible test results
	sort.Slice(f.SkipCRDs, func(i, j int) bool {
		return j > i
	})

	return &f, nil
}
