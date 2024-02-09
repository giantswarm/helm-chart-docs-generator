package git

import "github.com/giantswarm/microerror"

var CouldNotCloneRepositoryError = &microerror.Error{
	Kind: "CouldNotCloneRepositoryError",
	Desc: "The Git repository could not be cloned.",
}

// IsCouldNotCloneRepositoryError asserts CouldNotCloneRepositoryError
func IsCouldNotCloneRepositoryError(e error) bool {
	return microerror.Cause(e) == CouldNotCloneRepositoryError
}
