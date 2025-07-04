package linters

import (
	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

var _ Linter = new(SettingsAlphabetical)

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

	keys, ok := l.getSettingsKeys(golangci)
	if !ok {
		return nil
	}

	if _, ok := internal.IsAlphabetical(keys); !ok {
		reports = append(reports, internal.Report{
			Rule:    l.rule,
			Message: "linters.settings are not sorted alphabetically",
		})
	}

	return reports
}

func (l SettingsAlphabetical) Fix(golangci internal.Golangci) error {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}
	return linters.SortSettings()
}

func (l SettingsAlphabetical) getSettingsKeys(golangci internal.Golangci) ([]string, bool) {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil, false
	}

	settings, isSettings := linters.GetSettings()
	if !isSettings {
		return nil, false
	}

	return settings.GetKeys()
}
