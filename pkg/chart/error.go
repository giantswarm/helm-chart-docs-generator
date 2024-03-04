package chart

import "github.com/giantswarm/microerror"

var CouldNotGenerateChartFileError = &microerror.Error{
	Kind: "CouldNotGenerateChartFileError",
	Desc: "The Chart README file could not be generated.",
}

// IsCouldNotGenerateChartFileError asserts CouldNotGenerateChartFileError
func IsCouldNotGenerateChartFileError(e error) bool {
	return microerror.Cause(e) == CouldNotGenerateChartFileError
}

var CouldNotReadChartMetadataFileError = &microerror.Error{
	Kind: "CouldNotReadChartMetadataFileError",
	Desc: "The Chart README file could not be parsed.",
}

// IsCouldNotReadChartMetadataFileError asserts CouldNotReadChartMetadataFileError
func IsCouldNotReadChartMetadataFileError(e error) bool {
	return microerror.Cause(e) == CouldNotReadChartMetadataFileError
}

var CouldNotParsedChartFileError = &microerror.Error{
	Kind: "CouldNotParsedChartFileError",
	Desc: "The Chart README file could not be parsed.",
}

// IsCouldNotParsedChartFileError asserts CouldNotParsedChartFileError
func IsCouldNotParsedChartFileError(e error) bool {
	return microerror.Cause(e) == CouldNotParsedChartFileError
}
