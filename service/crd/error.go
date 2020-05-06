package crd

import "github.com/giantswarm/microerror"

var CouldNotReadCRDFileError = &microerror.Error{
	Kind: "CouldNotReadCRDFileError",
	Desc: "The CRD file could not be read.",
}

// IsCouldNotReadCRDFile asserts CouldNotReadCRDFileError
func IsCouldNotReadCRDFile(e error) bool {
	return microerror.Cause(e) == CouldNotReadCRDFileError
}

var CouldNotParseCRDFileError = &microerror.Error{
	Kind: "CouldNotParseCRDFileError",
	Desc: "The CRD file could not be parsed.",
}

// IsCouldNotParseCRDFile asserts CouldNotParseCRDFileError
func IsCouldNotParseCRDFile(e error) bool {
	return microerror.Cause(e) == CouldNotParseCRDFileError
}
