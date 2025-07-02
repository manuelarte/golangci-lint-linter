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

	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	{
		enable, hasEnable := linters.GetEnable()
		if hasEnable {
			if !internal.IsAlphabetical(enable) {
				reports = append(reports, internal.Report{
					Rule:    l.rule,
					Message: "linters.enable are not sorted alphabetically",
				})
			}
		}
	}
	{
		disable, hasDisable := linters.GetDisable()
		if hasDisable {
			if !internal.IsAlphabetical(disable) {
				reports = append(reports, internal.Report{
					Rule:    l.rule,
					Message: "linters.disable are not sorted alphabetically",
				})
			}
		}
	}

	return reports
}
