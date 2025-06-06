// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golang

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/format"

	"github.com/fatih/gomodifytags/modifytags"
	"github.com/block/ftl-golang-tools/gopls/internal/cache"
	"github.com/block/ftl-golang-tools/gopls/internal/cache/parsego"
	"github.com/block/ftl-golang-tools/gopls/internal/file"
	"github.com/block/ftl-golang-tools/gopls/internal/protocol"
	"github.com/block/ftl-golang-tools/gopls/internal/protocol/command"
	"github.com/block/ftl-golang-tools/gopls/internal/util/moreiters"
	internalastutil "github.com/block/ftl-golang-tools/internal/astutil"
	"github.com/block/ftl-golang-tools/internal/diff"
	"github.com/block/ftl-golang-tools/internal/tokeninternal"
)

// ModifyTags applies the given struct tag modifications to the specified struct.
func ModifyTags(ctx context.Context, snapshot *cache.Snapshot, fh file.Handle, args command.ModifyTagsArgs, m *modifytags.Modification) ([]protocol.DocumentChange, error) {
	pgf, err := snapshot.ParseGo(ctx, fh, parsego.Full)
	if err != nil {
		return nil, fmt.Errorf("error fetching package file: %v", err)
	}
	start, end, err := pgf.RangePos(args.Range)
	if err != nil {
		return nil, fmt.Errorf("error getting position information: %v", err)
	}
	// If the cursor is at a point and not a selection, we should use the entire enclosing struct.
	if start == end {
		cur, ok := pgf.Cursor.FindByPos(start, end)
		if !ok {
			return nil, fmt.Errorf("error finding start and end positions: %v", err)
		}
		curStruct, ok := moreiters.First(cur.Enclosing((*ast.StructType)(nil)))
		if !ok {
			return nil, fmt.Errorf("no enclosing struct type")
		}
		start, end = curStruct.Node().Pos(), curStruct.Node().End()
	}

	// Create a copy of the file node in order to avoid race conditions when we modify the node in Apply.
	cloned := internalastutil.CloneNode(pgf.File)
	fset := tokeninternal.FileSetFor(pgf.Tok)

	if err = m.Apply(fset, cloned, start, end); err != nil {
		return nil, fmt.Errorf("could not modify tags: %v", err)
	}

	// Construct a list of DocumentChanges based on the diff between the formatted node and the
	// original file content.
	var after bytes.Buffer
	if err := format.Node(&after, fset, cloned); err != nil {
		return nil, err
	}
	edits := diff.Bytes(pgf.Src, after.Bytes())
	if len(edits) == 0 {
		return nil, nil
	}
	textedits, err := protocol.EditsFromDiffEdits(pgf.Mapper, edits)
	if err != nil {
		return nil, fmt.Errorf("error computing edits for %s: %v", args.URI, err)
	}
	return []protocol.DocumentChange{
		protocol.DocumentChangeEdit(fh, textedits),
	}, nil
}
