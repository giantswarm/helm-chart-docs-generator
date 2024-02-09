package chart

import "github.com/giantswarm/microerror"

var CouldNotReadChartFileError = &microerror.Error{
	Kind: "CouldNotReadChartFileError",
	Desc: "The Chart README file could not be read.",
}

// IsCouldNotReadChartFileError asserts CouldNotReadChartFileError
func IsCouldNotReadChartFileError(e error) bool {
	return microerror.Cause(e) == CouldNotReadChartFileError
}

var CouldNotParsedChartFileError = &microerror.Error{
	Kind: "CouldNotParsedChartFileError",
	Desc: "The Chart README file could not be parsed.",
}

// IsCouldNotParsedChartFileError asserts CouldNotParsedChartFileError
func IsCouldNotParsedChartFileError(e error) bool {
	return microerror.Cause(e) == CouldNotParsedChartFileError
}
