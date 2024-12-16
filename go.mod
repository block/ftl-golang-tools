module github.com/block/ftl-golang-tools

go 1.22.0 // => default GODEBUG has gotypesalias=0

toolchain go1.23.3

require (
	github.com/google/go-cmp v0.6.0
	github.com/yuin/goldmark v1.4.13
	golang.org/x/exp v0.0.0-20241215155358-4a5509556b9e
	golang.org/x/mod v0.22.0
	golang.org/x/net v0.32.0
	golang.org/x/sync v0.10.0
	golang.org/x/telemetry v0.0.0-20240521205824-bda55230c457
	golang.org/x/tools v0.28.0
)

require golang.org/x/sys v0.28.0 // indirect
