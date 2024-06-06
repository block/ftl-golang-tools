// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/TBD54566975/golang-tools/gopls/internal/file"
	"github.com/TBD54566975/golang-tools/gopls/internal/golang"
	"github.com/TBD54566975/golang-tools/gopls/internal/label"
	"github.com/TBD54566975/golang-tools/gopls/internal/protocol"
	"github.com/TBD54566975/golang-tools/gopls/internal/telemetry"
	"github.com/TBD54566975/golang-tools/gopls/internal/template"
	"github.com/TBD54566975/golang-tools/internal/event"
)

func (s *server) References(ctx context.Context, params *protocol.ReferenceParams) (_ []protocol.Location, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("references")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.references", label.URI.Of(params.TextDocument.URI))
	defer done()

	fh, snapshot, release, err := s.fileOf(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	defer release()
	switch snapshot.FileKind(fh) {
	case file.Tmpl:
		return template.References(ctx, snapshot, fh, params)
	case file.Go:
		return golang.References(ctx, snapshot, fh, params.Position, params.Context.IncludeDeclaration)
	}
	return nil, nil // empty result
}
