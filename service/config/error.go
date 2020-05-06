package config

import "github.com/giantswarm/microerror"

var CouldNotReadConfigFileError = &microerror.Error{
	Kind: "CouldNotReadConfigFileError",
	Desc: "The configuration file could not be read.",
}

// IsCouldNotReadConfigFile asserts CouldNotReadConfigFileError
func IsCouldNotReadConfigFile(e error) bool {
	return microerror.Cause(e) == CouldNotReadConfigFileError
}

var CouldNotParseConfigFileError = &microerror.Error{
	Kind: "CouldNotParseConfigFileError",
	Desc: "The configuration file could not be parsed.",
}

// IsCouldNotParseConfigFile asserts CouldNotParseConfigFileError
func IsCouldNotParseConfigFile(e error) bool {
	return microerror.Cause(e) == CouldNotParseConfigFileError
}
