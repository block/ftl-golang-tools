// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unusedfunc_test

import (
	"path/filepath"
	"testing"

	"github.com/block/ftl-golang-tools/go/analysis/analysistest"
	"github.com/block/ftl-golang-tools/gopls/internal/analysis/unusedfunc"
	"github.com/block/ftl-golang-tools/internal/testfiles"
)

func Test(t *testing.T) {
	dir := testfiles.ExtractTxtarFileToTmp(t, filepath.Join(analysistest.TestData(), "basic.txtar"))
	analysistest.RunWithSuggestedFixes(t, dir, unusedfunc.Analyzer, "example.com/a")
}
