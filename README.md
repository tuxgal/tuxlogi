# zzzlogi

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tuxgal/zzzlogi)](https://pkg.go.dev/github.com/tuxgal/zzzlogi) [![Build](https://github.com/tuxgal/zzzlogi/actions/workflows/build.yml/badge.svg)](https://github.com/tuxgal/zzzlogi/actions/workflows/build.yml) [![Tests](https://github.com/tuxgal/zzzlogi/actions/workflows/tests.yml/badge.svg)](https://github.com/tuxgal/zzzlogi/actions/workflows/tests.yml) [![Lint](https://github.com/tuxgal/zzzlogi/actions/workflows/lint.yml/badge.svg)](https://github.com/tuxgal/zzzlogi/actions/workflows/lint.yml) [![CodeQL](https://github.com/tuxgal/zzzlogi/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/tuxgal/zzzlogi/actions/workflows/codeql-analysis.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/tuxgal/zzzlogi)](https://goreportcard.com/report/github.com/tuxgal/zzzlogi)

A go library that provides just a generic level logging interface.

The interface can be used to define an abstract logger in libraries or any
command line utilties, and allow users to customize the concrete logging
implementation to use.

This is currently used by
[`cablemodemutil`](https://github.com/tuxgal/cablemodemutil),
[`cablemodemcli`](https://github.com/tuxgal/cablemodemcli) and other projects
in `go` used by [`tuxgal`](https://github.com/tuxgal) and
[`tuxgalhomelab`](https://github.com/tuxgalhomelab).
