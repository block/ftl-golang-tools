// Package programmaticchecker provides a mechanism for running a set of analyzers on a package programmatically.
package programmaticchecker

import (
	"github.com/TBD54566975/golang-tools/go/analysis"
	"github.com/TBD54566975/golang-tools/go/analysis/internal/checker"
	"github.com/TBD54566975/golang-tools/go/packages"
)

// Config specifies the configuration for the programmatic checker.
type Config struct {
	// LoadConfig is the packages.Config to use when loading packages.
	LoadConfig packages.Config
	// ReverseImportExecutionOrder is true if packages that import a given package should execute _after_ the package itself.
	ReverseImportExecutionOrder bool
	// Patterns specify directory patterns for the package loader.
	Patterns []string
}

func Run(cfg Config, analyzers ...*analysis.Analyzer) (analyzerResults map[*analysis.Analyzer][]any, diagnostics []analysis.SimpleDiagnostic, err error) {
	if err := analysis.Validate(analyzers); err != nil {
		return nil, nil, err
	}

	return checker.RunWithResult(cfg.Patterns, analyzers,
		checker.WithLoadConfig(cfg.LoadConfig),
		checker.WithReverseImportExecutionOrder(cfg.ReverseImportExecutionOrder),
	)
}
