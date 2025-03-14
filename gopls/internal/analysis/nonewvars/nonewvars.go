// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nonewvars defines an Analyzer that applies suggested fixes
// to errors of the type "no new variables on left side of :=".
package nonewvars

import (
	"bytes"
	_ "embed"
	"go/ast"
	"go/format"
	"go/token"

	"github.com/block/ftl-golang-tools/go/analysis"
	"github.com/block/ftl-golang-tools/go/analysis/passes/inspect"
	"github.com/block/ftl-golang-tools/go/ast/inspector"
	"github.com/block/ftl-golang-tools/internal/analysisinternal"
)

//go:embed doc.go
var doc string

var Analyzer = &analysis.Analyzer{
	Name:             "nonewvars",
	Doc:              analysisinternal.MustExtractDoc(doc, "nonewvars"),
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	Run:              run,
	RunDespiteErrors: true,
	URL:              "https://pkg.go.dev/github.com/block/ftl-golang-tools/gopls/internal/analysis/nonewvars",
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if len(pass.TypeErrors) == 0 {
		return nil, nil
	}

	nodeFilter := []ast.Node{(*ast.AssignStmt)(nil)}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		assignStmt, _ := n.(*ast.AssignStmt)
		// We only care about ":=".
		if assignStmt.Tok != token.DEFINE {
			return
		}

		var file *ast.File
		for _, f := range pass.Files {
			if f.FileStart <= assignStmt.Pos() && assignStmt.Pos() < f.FileEnd {
				file = f
				break
			}
		}
		if file == nil {
			return
		}

		for _, err := range pass.TypeErrors {
			if !FixesError(err.Msg) {
				continue
			}
			if assignStmt.Pos() > err.Pos || err.Pos >= assignStmt.End() {
				continue
			}
			var buf bytes.Buffer
			if err := format.Node(&buf, pass.Fset, file); err != nil {
				continue
			}
			pass.Report(analysis.Diagnostic{
				Pos:     err.Pos,
				End:     analysisinternal.TypeErrorEndPos(pass.Fset, buf.Bytes(), err.Pos),
				Message: err.Msg,
				SuggestedFixes: []analysis.SuggestedFix{{
					Message: "Change ':=' to '='",
					TextEdits: []analysis.TextEdit{{
						Pos: err.Pos,
						End: err.Pos + 1,
					}},
				}},
			})
		}
	})
	return nil, nil
}

func FixesError(msg string) bool {
	return msg == "no new variables on left side of :="
}
