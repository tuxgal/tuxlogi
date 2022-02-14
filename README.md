# zzzlogi

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tuxdude/zzzlogi)](https://pkg.go.dev/github.com/tuxdude/zzzlogi) [![Build](https://github.com/Tuxdude/zzzlogi/actions/workflows/build.yml/badge.svg)](https://github.com/Tuxdude/zzzlogi/actions/workflows/build.yml) [![Tests](https://github.com/Tuxdude/zzzlogi/actions/workflows/tests.yml/badge.svg)](https://github.com/Tuxdude/zzzlogi/actions/workflows/tests.yml) [![Lint](https://github.com/Tuxdude/zzzlogi/actions/workflows/lint.yml/badge.svg)](https://github.com/Tuxdude/zzzlogi/actions/workflows/lint.yml) [![CodeQL](https://github.com/Tuxdude/zzzlogi/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/Tuxdude/zzzlogi/actions/workflows/codeql-analysis.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/tuxdude/zzzlogi)](https://goreportcard.com/report/github.com/tuxdude/zzzlogi)

A go library that provides just a generic level logging interface.

The interface can be used to define an abstract logger in libraries or any
command line utilties, and allow users to customize the concrete logging
implementation to use.

This is currently used by
[`cablemodemutil`](https://github.com/Tuxdude/cablemodemutil),
[`cablemodemcli`](https://github.com/Tuxdude/cablemodemcli) and other projects
in `go` used by [`Tuxdude`](https://github.com/Tuxdude) and
[`TuxdudeHomeLab`](https://github.com/TuxdudeHomeLab).
