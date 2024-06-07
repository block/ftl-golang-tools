// Package programmaticchecker provides a mechanism for running a set of analyzers on a package programmatically.
package programmaticchecker

import (
	"github.com/TBD54566975/golang-tools/go/analysis"
	"github.com/TBD54566975/golang-tools/go/analysis/internal/checker"
	"github.com/TBD54566975/golang-tools/go/packages"
)

type AnalyzerResults map[*analysis.Analyzer]any

var globalResults map[*analysis.Analyzer]interface{}

var resultFetcher = &analysis.Analyzer{
	Name: "resultFetcher",
	Doc:  "propogates the results from all analyzers to return to the caller",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		globalResults = pass.ResultOf
		return nil, nil
	},
}

func Run(patterns []string, loadConfig packages.Config, analyzers ...*analysis.Analyzer) (*AnalyzerResults, error) {
	resultFetcher.Requires = analyzers
	withResult := append(analyzers, resultFetcher)
	if err := analysis.Validate(withResult); err != nil {
		return nil, err
	}

	checker.Run(patterns, withResult, checker.WithLoadConfig(loadConfig))

	result := make(AnalyzerResults)
	for _, a := range analyzers {
		result[a] = globalResults[a]
	}
	return &result, nil
}
