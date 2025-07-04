package linters

import (
	"cmp"
	"github.com/manuelarte/golangci-lint-linter-go/internal"
	"slices"
)

var _ Linter = new(LinterFieldsSorted)

type LinterFieldsSorted struct {
	rule          internal.RuleCode
	expectedOrder map[string]int
}

func NewLinterFieldsSorted() *LinterFieldsSorted {
	expectedOrder := map[string]int{"default": 0, "enable": 1, "disable": 2, "settings": 3, "exclusions": 4}
	return &LinterFieldsSorted{
		rule:          internal.GC013,
		expectedOrder: expectedOrder,
	}
}

func (l LinterFieldsSorted) Lint(golangci internal.Golangci) []internal.Report {
	reports := make([]internal.Report, 0)

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
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters fields are not sorted",
		})
	}

	return reports
}

func (l LinterFieldsSorted) Fix(golangci internal.Golangci) error {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	return linters.SortFields(l.expectedOrder)
}
