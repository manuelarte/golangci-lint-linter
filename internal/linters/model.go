package linters

import "github.com/manuelarte/golangci-lint-linter-go/internal"

type Linter interface {
	Lint(golangci internal.Golangci) []internal.Report
}

type Fixer interface {
	Fix(golangci internal.Golangci) error
}
