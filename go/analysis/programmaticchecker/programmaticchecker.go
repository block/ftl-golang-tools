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

type Config struct {
	// LoadConfig is the packages.Config to use when loading packages.
	LoadConfig packages.Config
	// RunDespiteLoadErrors specifies whether to run the analysis even if there are package load errors.
	RunDespiteLoadErrors bool
	// Patterns specify directory patterns for the package loader.
	Patterns []string
}

func Run(cfg Config, analyzers ...*analysis.Analyzer) (*AnalyzerResults, error) {
	resultFetcher.Requires = analyzers
	resultFetcher.RunDespiteErrors = cfg.RunDespiteLoadErrors
	withResult := append(analyzers, resultFetcher)
	if err := analysis.Validate(withResult); err != nil {
		return nil, err
	}

	checker.Run(cfg.Patterns, withResult, checker.WithLoadConfig(cfg.LoadConfig), checker.WithRunDespiteLoadErrors(cfg.RunDespiteLoadErrors))

	result := make(AnalyzerResults)
	for _, a := range analyzers {
		if r, ok := globalResults[a]; ok {
			result[a] = r
		}
	}
	return &result, nil
}
