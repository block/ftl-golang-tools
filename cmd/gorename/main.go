// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gorename command performs precise type-safe renaming of
// identifiers in Go source code.
//
// Run with -help for usage information, or view the Usage constant in
// package github.com/block/ftl-golang-tools/refactor/rename, which contains most of
// the implementation.
package main // import "github.com/block/ftl-golang-tools/cmd/gorename"

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"

	"github.com/block/ftl-golang-tools/go/buildutil"
	"github.com/block/ftl-golang-tools/refactor/rename"
)

var (
	offsetFlag = flag.String("offset", "", "file and byte offset of identifier to be renamed, e.g. 'file.go:#123'.  For use by editors.")
	fromFlag   = flag.String("from", "", "identifier to be renamed; see -help for formats")
	toFlag     = flag.String("to", "", "new name for identifier")
	helpFlag   = flag.Bool("help", false, "show usage message")
)

func init() {
	flag.Var((*buildutil.TagsFlag)(&build.Default.BuildTags), "tags", buildutil.TagsFlagDoc)
	flag.BoolVar(&rename.Force, "force", false, "proceed, even if conflicts were reported")
	flag.BoolVar(&rename.Verbose, "v", false, "print verbose information")
	flag.BoolVar(&rename.Diff, "d", false, "display diffs instead of rewriting files")
	flag.StringVar(&rename.DiffCmd, "diffcmd", "diff", "diff command invoked when using -d")
}

func main() {
	log.SetPrefix("gorename: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	if *helpFlag || (*offsetFlag == "" && *fromFlag == "" && *toFlag == "") {
		fmt.Print(rename.Usage)
		return
	}

	if err := rename.Main(&build.Default, *offsetFlag, *fromFlag, *toFlag); err != nil {
		if err != rename.ConflictError {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
