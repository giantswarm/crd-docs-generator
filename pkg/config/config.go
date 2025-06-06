package config

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/giantswarm/microerror"
	"github.com/goccy/go-yaml"
)

// FromFile represent a config file content.
type FromFile struct {
	SourceRepositories []SourceRepository `yaml:"source_repositories"`
	TemplatePath       string             `yaml:"template_path"`
	OutputPath         string             `yaml:"output_path"`
}

// SourceRepository has details about a
// source repository to find CRDs in.
type SourceRepository struct {
	URL             string             `yaml:"url"`
	Organization    string             `yaml:"organization"`
	ShortName       string             `yaml:"short_name"`
	Metadata        map[string]CRDItem `yaml:"metadata"`
	CommitReference string             `yaml:"commit_reference"`
	AnnotationsPath []string           `yaml:"annotations_paths"`
	CRDPaths        []string           `yaml:"crd_paths"`
	CRPaths         []string           `yaml:"cr_paths"`
}

type CRDItem struct {
	Owners      []string     `yaml:"owner,omitempty"`
	Topics      []string     `yaml:"topics,omitempty"`
	Providers   []string     `yaml:"provider,omitempty"`
	Hidden      bool         `yaml:"hidden,omitempty"`
	Deprecation *Deprecation `yaml:"deprecation,omitempty"`
}

type Deprecation struct {
	Info       string                 `yaml:"info,omitempty"`
	ReplacedBy *DeprecationReplacedBy `yaml:"replaced_by,omitempty"`
}

type DeprecationReplacedBy struct {
	FullName  string `yaml:"full_name"`
	ShortName string `yaml:"short_name"`
}

// Read reads a config file and returns a struct.
func Read(path string) (*FromFile, error) {
	f := &FromFile{}

	path = filepath.Clean(path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, microerror.Maskf(CouldNotReadConfigFileError, "%s", err.Error())
	}

	reader := bytes.NewReader(data)
	// Fail on unknown fields.
	decoder := yaml.NewDecoder(reader, yaml.DisallowUnknownField())
	err = decoder.Decode(f)
	if err != nil {
		return nil, microerror.Maskf(CouldNotParseConfigFileError, "%s", err.Error())
	}

	return f, nil
}
