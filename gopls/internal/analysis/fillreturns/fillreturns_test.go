// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fillreturns_test

import (
	"testing"

	"github.com/block/ftl-golang-tools/go/analysis/analysistest"
	"github.com/block/ftl-golang-tools/gopls/internal/analysis/fillreturns"
	"github.com/block/ftl-golang-tools/internal/aliases"
)

func Test(t *testing.T) {
	// TODO(golang/go#65294): update expectations and delete this
	// check once we update go.mod to go1.23 so that
	// gotypesalias=1 becomes the only supported mode.
	if aliases.Enabled() {
		t.Skip("expectations need updating for materialized aliases")
	}

	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, fillreturns.Analyzer, "a", "typeparams")
}
