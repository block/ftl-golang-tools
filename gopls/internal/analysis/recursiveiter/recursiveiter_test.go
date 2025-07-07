// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recursiveiter_test

import (
	"testing"

	"github.com/block/ftl-golang-tools/go/analysis/analysistest"
	"github.com/block/ftl-golang-tools/gopls/internal/analysis/recursiveiter"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, recursiveiter.Analyzer, "a")
}
