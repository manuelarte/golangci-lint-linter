package linters

import (
	"cmp"
	"slices"

	"github.com/manuelarte/golangci-lint-linter/internal"
	"github.com/manuelarte/golangci-lint-linter/models"
)

var _ Linter = new(LinterFieldsSorted)

type LinterFieldsSorted struct {
	rule          models.RuleCode
	expectedOrder map[string]int
}

func NewLinterFieldsSorted() *LinterFieldsSorted {
	expectedOrder := map[string]int{"default": 0, "enable": 1, "disable": 2, "settings": 3, "exclusions": 4}

	return &LinterFieldsSorted{
		rule:          models.GC013,
		expectedOrder: expectedOrder,
	}
}

func (l LinterFieldsSorted) Rule() models.RuleCode {
	return l.rule
}

func (l LinterFieldsSorted) Lint(golangci models.Golangci) []models.Report {
	reports := make([]models.Report, 0)

	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	fields, ok := linters.GetFieldNames()
	if !ok {
		return nil
	}

	sorted := slices.Clone(fields)
	slices.SortFunc(sorted, func(a, b string) int {
		aIndex := l.expectedOrder[a]
		bIndex := l.expectedOrder[b]

		return cmp.Compare(aIndex, bIndex)
	})

	if !internal.EqualArray(fields, sorted) {
		reports = append(reports, models.Report{
			Rule:    l.rule,
			Message: "linters fields are not sorted",
		})
	}

	return reports
}

func (l LinterFieldsSorted) Fix(golangci models.Golangci) error {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	return linters.SortFields(l.expectedOrder)
}
