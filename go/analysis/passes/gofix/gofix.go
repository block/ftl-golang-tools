// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gofix defines an analyzer that checks go:fix directives.
package gofix

import (
	_ "embed"

	"github.com/block/ftl-golang-tools/go/analysis"
	"github.com/block/ftl-golang-tools/go/analysis/passes/inspect"
	"github.com/block/ftl-golang-tools/go/ast/inspector"
	"github.com/block/ftl-golang-tools/internal/analysisinternal"
	"github.com/block/ftl-golang-tools/internal/gofix/findgofix"
)

//go:embed doc.go
var doc string

var Analyzer = &analysis.Analyzer{
	Name:     "gofixdirective",
	Doc:      analysisinternal.MustExtractDoc(doc, "gofixdirective"),
	URL:      "https://pkg.go.dev/github.com/block/ftl-golang-tools/go/analysis/passes/gofix",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	root := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector).Root()
	findgofix.Find(pass, root, nil)
	return nil, nil
}
