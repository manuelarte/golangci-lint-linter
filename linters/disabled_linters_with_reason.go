package linters

import (
	"fmt"

	"github.com/manuelarte/golangci-lint-linter/models"
)

var _ Linter = new(DisabledLintersWithReason)

type DisabledLintersWithReason struct {
	rule models.RuleCode
}

func NewDisabledLintersWithReason() *DisabledLintersWithReason {
	return &DisabledLintersWithReason{
		rule: models.GC021,
	}
}

func (l DisabledLintersWithReason) Rule() models.RuleCode {
	return l.rule
}

func (l DisabledLintersWithReason) Lint(golangci models.Golangci) []models.Report {
	reports := make([]models.Report, 0)

	linters, ok := golangci.GetLinters()
	if !ok {
		return nil
	}

	disabled, isOk := linters.GetDisable()
	if !isOk {
		return nil
	}

	for i, disable := range disabled {
		_, hasComment := golangci.GetComment(fmt.Sprintf("$.linters.disable[%d]", i))
		if !hasComment {
			reports = append(reports, models.Report{
				Rule:    l.rule,
				Message: fmt.Sprintf("linters.disable.%s does not have a reason", disable),
			})

			break
		}
	}

	return reports
}
