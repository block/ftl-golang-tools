// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gopls (pronounced “go please”) is an LSP server for Go.
// The Language Server Protocol allows any text editor
// to be extended with IDE-like features;
// see https://langserver.org/ for details.
//
// See https://github.com/golang/tools/blob/master/gopls/README.md
// for the most up-to-date documentation.
package main // import "github.com/block/ftl-golang-tools/gopls"

import (
	"context"
	"os"

	"golang.org/x/telemetry"
	"github.com/block/ftl-golang-tools/gopls/internal/cmd"
	versionpkg "github.com/block/ftl-golang-tools/gopls/internal/version"
	"github.com/block/ftl-golang-tools/internal/tool"
)

var version = "" // if set by the linker, overrides the gopls version

func main() {
	versionpkg.VersionOverride = version

	telemetry.Start(telemetry.Config{
		ReportCrashes: true,
		Upload:        true,
	})

	ctx := context.Background()
	tool.Main(ctx, cmd.New(), os.Args[1:])
}
