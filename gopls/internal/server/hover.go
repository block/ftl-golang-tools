// Copyright 2019 The Go Authors. All rights reserved.
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
	"github.com/TBD54566975/x/tools/gopls/internal/telemetry"
	"github.com/TBD54566975/x/tools/gopls/internal/template"
	"github.com/TBD54566975/x/tools/gopls/internal/work"
	"github.com/TBD54566975/x/tools/internal/event"
)

func (s *server) Hover(ctx context.Context, params *protocol.HoverParams) (_ *protocol.Hover, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("hover")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.hover", label.URI.Of(params.TextDocument.URI))
	defer done()

	fh, snapshot, release, err := s.fileOf(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	defer release()

	switch snapshot.FileKind(fh) {
	case file.Mod:
		return mod.Hover(ctx, snapshot, fh, params.Position)
	case file.Go:
		return golang.Hover(ctx, snapshot, fh, params.Position)
	case file.Tmpl:
		return template.Hover(ctx, snapshot, fh, params.Position)
	case file.Work:
		return work.Hover(ctx, snapshot, fh, params.Position)
	}
	return nil, nil // empty result
}
