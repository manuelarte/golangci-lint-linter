package linters

import (
	"slices"

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

	return reports
}

func isAlphabetical(original []string) bool {
	sorted := slices.Clone(original)
	slices.Sort(sorted)

	return internal.EqualArray(original, sorted)
}
