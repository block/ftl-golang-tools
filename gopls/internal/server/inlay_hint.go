// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/TBD54566975/x/tools/gopls/internal/file"
	"github.com/TBD54566975/x/tools/gopls/internal/golang"
	"github.com/TBD54566975/x/tools/gopls/internal/label"
	"github.com/TBD54566975/x/tools/gopls/internal/mod"
	"github.com/TBD54566975/x/tools/gopls/internal/protocol"
	"github.com/TBD54566975/x/tools/internal/event"
)

func (s *server) InlayHint(ctx context.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	ctx, done := event.Start(ctx, "lsp.Server.inlayHint", label.URI.Of(params.TextDocument.URI))
	defer done()

	fh, snapshot, release, err := s.fileOf(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	defer release()

	switch snapshot.FileKind(fh) {
	case file.Mod:
		return mod.InlayHint(ctx, snapshot, fh, params.Range)
	case file.Go:
		return golang.InlayHint(ctx, snapshot, fh, params.Range)
	}
	return nil, nil // empty result
}
