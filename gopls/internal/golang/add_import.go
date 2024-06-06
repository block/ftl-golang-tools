// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golang

import (
	"context"

	"github.com/TBD54566975/golang-tools/gopls/internal/cache"
	"github.com/TBD54566975/golang-tools/gopls/internal/cache/parsego"
	"github.com/TBD54566975/golang-tools/gopls/internal/file"
	"github.com/TBD54566975/golang-tools/gopls/internal/protocol"
	"github.com/TBD54566975/golang-tools/internal/imports"
)

// AddImport adds a single import statement to the given file
func AddImport(ctx context.Context, snapshot *cache.Snapshot, fh file.Handle, importPath string) ([]protocol.TextEdit, error) {
	pgf, err := snapshot.ParseGo(ctx, fh, parsego.Full)
	if err != nil {
		return nil, err
	}
	return ComputeOneImportFixEdits(snapshot, pgf, &imports.ImportFix{
		StmtInfo: imports.ImportInfo{
			ImportPath: importPath,
		},
		FixType: imports.AddImport,
	})
}
