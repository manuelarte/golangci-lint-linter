package linters

import (
	"github.com/manuelarte/golangci-lint-linter/internal"
	"github.com/manuelarte/golangci-lint-linter/models"
)

type Linter interface {
	Lint(golangci models.Golangci) []internal.Report
}

type Fixer interface {
	Fix(golangci models.Golangci) error
}
