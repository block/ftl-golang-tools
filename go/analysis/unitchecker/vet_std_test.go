// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unitchecker_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/block/ftl-golang-tools/go/analysis/passes/appends"
	"github.com/block/ftl-golang-tools/go/analysis/passes/asmdecl"
	"github.com/block/ftl-golang-tools/go/analysis/passes/assign"
	"github.com/block/ftl-golang-tools/go/analysis/passes/atomic"
	"github.com/block/ftl-golang-tools/go/analysis/passes/bools"
	"github.com/block/ftl-golang-tools/go/analysis/passes/buildtag"
	"github.com/block/ftl-golang-tools/go/analysis/passes/cgocall"
	"github.com/block/ftl-golang-tools/go/analysis/passes/composite"
	"github.com/block/ftl-golang-tools/go/analysis/passes/copylock"
	"github.com/block/ftl-golang-tools/go/analysis/passes/defers"
	"github.com/block/ftl-golang-tools/go/analysis/passes/directive"
	"github.com/block/ftl-golang-tools/go/analysis/passes/errorsas"
	"github.com/block/ftl-golang-tools/go/analysis/passes/framepointer"
	"github.com/block/ftl-golang-tools/go/analysis/passes/httpresponse"
	"github.com/block/ftl-golang-tools/go/analysis/passes/ifaceassert"
	"github.com/block/ftl-golang-tools/go/analysis/passes/loopclosure"
	"github.com/block/ftl-golang-tools/go/analysis/passes/lostcancel"
	"github.com/block/ftl-golang-tools/go/analysis/passes/nilfunc"
	"github.com/block/ftl-golang-tools/go/analysis/passes/printf"
	"github.com/block/ftl-golang-tools/go/analysis/passes/shift"
	"github.com/block/ftl-golang-tools/go/analysis/passes/sigchanyzer"
	"github.com/block/ftl-golang-tools/go/analysis/passes/stdmethods"
	"github.com/block/ftl-golang-tools/go/analysis/passes/stdversion"
	"github.com/block/ftl-golang-tools/go/analysis/passes/stringintconv"
	"github.com/block/ftl-golang-tools/go/analysis/passes/structtag"
	"github.com/block/ftl-golang-tools/go/analysis/passes/testinggoroutine"
	"github.com/block/ftl-golang-tools/go/analysis/passes/tests"
	"github.com/block/ftl-golang-tools/go/analysis/passes/timeformat"
	"github.com/block/ftl-golang-tools/go/analysis/passes/unmarshal"
	"github.com/block/ftl-golang-tools/go/analysis/passes/unreachable"
	"github.com/block/ftl-golang-tools/go/analysis/passes/unusedresult"
	"github.com/block/ftl-golang-tools/go/analysis/unitchecker"
)

// vet is the entrypoint of this executable when ENTRYPOINT=vet.
// Keep consistent with the actual vet in GOROOT/src/cmd/vet/main.go.
func vet() {
	unitchecker.Main(
		appends.Analyzer,
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		defers.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stdversion.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		testinggoroutine.Analyzer,
		tests.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		// unsafeptr.Analyzer, // currently reports findings in runtime
		unusedresult.Analyzer,
	)
}

// TestVetStdlib runs the same analyzers as the actual vet over the
// standard library, using go vet and unitchecker, to ensure that
// there are no findings.
func TestVetStdlib(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in -short mode")
	}
	if version := runtime.Version(); !strings.HasPrefix(version, "devel") {
		t.Skipf("This test is only wanted on development branches where code can be easily fixed. Skipping because runtime.Version=%q.", version)
	}

	cmd := exec.Command("go", "vet", "-vettool="+os.Args[0], "std")
	cmd.Env = append(os.Environ(), "ENTRYPOINT=vet")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("go vet std failed (%v):\n%s", err, out)
	}
}
