// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// The unusedfunc command runs the unusedfunc analyzer.
package main

import (
	"github.com/block/ftl-golang-tools/go/analysis/singlechecker"
	"github.com/block/ftl-golang-tools/gopls/internal/analysis/unusedfunc"
)

func main() { singlechecker.Main(unusedfunc.Analyzer) }
