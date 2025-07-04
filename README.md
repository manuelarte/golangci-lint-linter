# Golangci-lint Linter

[![Go](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/go.yml)
[![PR checks](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/pr-checks.yml/badge.svg)](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/pr-checks.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/manuelarte/golangci-lint-linter)](https://goreportcard.com/report/github.com/manuelarte/golangci-lint-linter)
![version](https://img.shields.io/github/v/release/manuelarte/golangci-lint-linter)

Linter for the configuration files for the well known linter [`golangci-lint`][golangci-lint].

## ⬇️  Getting Started

To install it run:

```bash
go install github.com/manuelarte/golangci-lint-linter@latest
```

To run it:

```bash
golangci-lint-linter [FILE] [--fix]
```

- `FILE`: the path to the `.golangci.yml` to be linter.
- `fix`: update file to resolve fixable issues. Default `false`.

## 📜 Rules

The list of linting rules can be found in [RULES.md](./RULES.md).

[golangci-lint]: https://golangci-lint.run/
