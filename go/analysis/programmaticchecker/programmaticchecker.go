// Package programmaticchecker provides a mechanism for running a set of analyzers on a package programmatically.
package programmaticchecker

import (
	"sync"

	"github.com/TBD54566975/golang-tools/go/analysis"
	"github.com/TBD54566975/golang-tools/go/analysis/internal/checker"
	"github.com/TBD54566975/golang-tools/go/packages"
)

type AnalyzerResults map[*analysis.Analyzer]any

var globalResults sync.Map

var resultFetcher = &analysis.Analyzer{
	Name: "resultFetcher",
	Doc:  "propogates the results from all analyzers to return to the caller",
	Run: func(pass *analysis.Pass) (interface{}, error) {
		for k, v := range pass.ResultOf {
			globalResults.Store(k, v)
		}
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

func Run(cfg Config, analyzers ...*analysis.Analyzer) (AnalyzerResults, error) {
	resultFetcher.Requires = analyzers
	resultFetcher.RunDespiteErrors = cfg.RunDespiteLoadErrors
	withResult := append(analyzers, resultFetcher)
	if err := analysis.Validate(withResult); err != nil {
		return nil, err
	}

	checker.Run(cfg.Patterns, withResult, checker.WithLoadConfig(cfg.LoadConfig), checker.WithRunDespiteLoadErrors(cfg.RunDespiteLoadErrors))

	result := AnalyzerResults{}
	globalResults.Range(func(key, value interface{}) bool {
		analyzer, ok := key.(*analysis.Analyzer)
		if !ok {
			return false
		}
		result[analyzer] = value
		return true
	})
	return result, nil
}
