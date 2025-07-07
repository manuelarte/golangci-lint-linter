package linters

import (
	"github.com/manuelarte/golangci-lint-linter/internal"
	"github.com/manuelarte/golangci-lint-linter/models"
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

func (l SettingsAlphabetical) Lint(golangci models.Golangci) []internal.Report {
	reports := make([]internal.Report, 0)

	keys, hasKeys := l.getSettingsKeys(golangci)
	if !hasKeys {
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

func (l SettingsAlphabetical) Fix(golangci models.Golangci) error {
	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	return linters.SortSettings()
}

func (l SettingsAlphabetical) getSettingsKeys(golangci models.Golangci) ([]string, bool) {
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
