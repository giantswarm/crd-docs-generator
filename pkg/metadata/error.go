package metadata

import "github.com/giantswarm/microerror"

var CouldNotReadFileError = &microerror.Error{
	Kind: "CouldNotReadFileError",
	Desc: "The metadata file could not be read.",
}

// IsCouldNotReadFile asserts CouldNotReadFileError
func IsCouldNotReadFile(e error) bool {
	return microerror.Cause(e) == CouldNotReadFileError
}

var CouldNotParseFileError = &microerror.Error{
	Kind: "CouldNotParseFileError",
	Desc: "The metadata file could not be parsed.",
}

// IsCouldNotParseFile asserts CouldNotParseFileError
func IsCouldNotParseFile(e error) bool {
	return microerror.Cause(e) == CouldNotParseFileError
}
