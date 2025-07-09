package linters

import (
	"github.com/manuelarte/golangci-lint-linter/models"
)

type Linter interface {
	Rule() models.RuleCode
	Lint(golangci models.Golangci) []models.Report
}

type Fixer interface {
	Fix(golangci models.Golangci) error
}
