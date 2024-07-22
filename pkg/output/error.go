package output

import "github.com/giantswarm/microerror"

var cannotOpenTemplate = &microerror.Error{
	Kind: "cannotOpenTemplate",
}

var cannotUnmarshalCRD = &microerror.Error{
	Kind: "cannotUnmarshalCRD",
}
