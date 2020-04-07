package error

import (
	"github.com/giantswarm/microerror"
)

var ExecutionError = &microerror.Error{
	Kind: "executionError",
}
