// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.20

package analyzer_test

import (
	"testing"

	"github.com/block/ftl-golang-tools/go/analysis/analysistest"
	inlineanalyzer "github.com/block/ftl-golang-tools/internal/refactor/inline/analyzer"
)

func TestAnalyzer(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), inlineanalyzer.Analyzer, "a", "b")
}
