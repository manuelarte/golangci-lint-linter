package linters

import (
	"fmt"

	"github.com/manuelarte/golangci-lint-linter/internal"
)

var _ Linter = new(DisabledLintersWithReason)

type DisabledLintersWithReason struct {
	rule internal.RuleCode
}

func NewDisabledLintersWithReason() *DisabledLintersWithReason {
	return &DisabledLintersWithReason{
		rule: internal.GC021,
	}
}

func (l DisabledLintersWithReason) Lint(golangci internal.Golangci) []internal.Report {
	reports := make([]internal.Report, 0)

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
			reports = append(reports, internal.Report{
				Rule:    l.rule,
				Message: fmt.Sprintf("linters.disable.%s does not have a reason", disable),
			})

			break
		}
	}

	return reports
}
