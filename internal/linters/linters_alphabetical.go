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
	if !internal.IsAlphabetical(enable) {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.enable are not sorted alphabetically",
		})
	}

	if !internal.IsAlphabetical(disable) {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.disable are not sorted alphabetically",
		})
	}

	return reports
}

func (l LintersAlphabetical) Fix(golangci internal.Golangci) {
	// linters, ok := golangci.GetLinters()
}

func (l LintersAlphabetical) getEnableAndDisable(golangci internal.Golangci) (enable []string, disable []string) {
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
