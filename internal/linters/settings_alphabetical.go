package linters

import (
	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

var _ Linter = new(SettingsAlphabetical)

//nolint:revive // is checking alphabetical order for the linters
type SettingsAlphabetical struct {
	rule internal.RuleCode
}

func NewSettingsAlphabetical() *SettingsAlphabetical {
	return &SettingsAlphabetical{
		rule: internal.GC002,
	}
}

func (l SettingsAlphabetical) Lint(golangci internal.Golangci) []internal.Report {
	reports := make([]internal.Report, 0)

	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	settings, isSettings := linters.GetSettings()
	if !isSettings {
		return nil
	}

	keys, ok := settings.GetKeys()
	if !ok {
		return nil
	}

	if !internal.IsAlphabetical(keys) {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.settings are not sorted alphabetically",
		})
	}

	return reports
}
