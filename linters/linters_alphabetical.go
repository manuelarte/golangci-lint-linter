package linters

import (
	"github.com/manuelarte/golangci-lint-linter/internal"
	"github.com/manuelarte/golangci-lint-linter/models"
)

var _ Linter = new(LintersAlphabetical)

//nolint:revive // is checking alphabetical order for the linters
type LintersAlphabetical struct {
	rule models.RuleCode
}

func NewLintersAlphabetical() *LintersAlphabetical {
	return &LintersAlphabetical{
		rule: models.GC001,
	}
}

func (l LintersAlphabetical) Rule() models.RuleCode {
	return l.rule
}

func (l LintersAlphabetical) Lint(golangci models.Golangci) []models.Report {
	reports := make([]models.Report, 0)

	enable, disable := l.getEnableAndDisable(golangci)
	if _, ok := internal.IsAlphabetical(enable); !ok {
		reports = append(reports, models.Report{
			Rule:    l.rule,
			Message: "linters.enable are not sorted alphabetically",
		})
	}

	if _, ok := internal.IsAlphabetical(disable); !ok {
		reports = append(reports, models.Report{
			Rule:    l.rule,
			Message: "linters.disable are not sorted alphabetically",
		})
	}

	return reports
}

func (l LintersAlphabetical) Fix(golangci models.Golangci) error {
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
func (l LintersAlphabetical) getEnableAndDisable(golangci models.Golangci) (enable, disable []string) {
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
