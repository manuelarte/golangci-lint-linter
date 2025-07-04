package linters

import (
	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

var _ Linter = new(LintersAlphabetical)

//nolint:revive // is checking alphabetical order for the linters
type LintersAlphabetical struct {
	rule internal.RuleCode
}

func NewLintersAlphabetical() *LintersAlphabetical {
	return &LintersAlphabetical{
		rule: internal.GC001,
	}
}

func (l LintersAlphabetical) Lint(golangci internal.Golangci) []internal.Report {
	reports := make([]internal.Report, 0)

	enable, disable := l.getEnableAndDisable(golangci)
	if _, ok := internal.IsAlphabetical(enable); !ok {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.enable are not sorted alphabetically",
		})
	}

	if _, ok := internal.IsAlphabetical(disable); !ok {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.disable are not sorted alphabetically",
		})
	}

	return reports
}

func (l LintersAlphabetical) Fix(golangci internal.Golangci) error {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	err := linters.SortEnable()
	if err != nil {
		return err
	}

	err = linters.SortDisable()
	if err != nil {
		return err
	}

	return nil
}

//nolint:nonamedreturns // two return same type
func (l LintersAlphabetical) getEnableAndDisable(golangci internal.Golangci) (enable, disable []string) {
	linters, ok := golangci.GetLinters()
	if !ok {
		return
	}

	e, hasEnable := linters.GetEnable()
	if hasEnable {
		enable = e
	}

	d, hasDisable := linters.GetDisable()
	if hasDisable {
		disable = d
	}

	return
}
