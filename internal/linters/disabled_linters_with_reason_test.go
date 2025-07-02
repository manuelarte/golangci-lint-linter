package linters

import (
	"testing"

	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

func TestDisabledLintersWithReason_Lint(t *testing.T) {
	testCases := map[string]struct {
		input       []byte
		expectedLen int
	}{
		"linters field not present": {
			input: []byte(`
version: 2
`),
			expectedLen: 0,
		},
		"none linters disable have a reason": {
			input: []byte(`
version: 2
linters:
  disable:
    - tagliatelle
    - funcorder
`),
			expectedLen: 1,
		},
		"some linters disable have a reason": {
			input: []byte(`
version: 2
linters:
  disable:
    - tagliatelle # to be checked in code reviews.
    - funcorder
`),
			expectedLen: 1,
		},
		"all linters disable have a reason": {
			input: []byte(`
version: 2
linters:
  disable:
    - tagliatelle # json tags to be checked in code reviews.
    - funcorder # function order to be checked in code reviews,
`),
			expectedLen: 0,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			golangci, err := internal.Parse(test.input)
			if err != nil {
				t.Fatal(err)
			}

			lint := NewDisabledLintersWithReason()

			reports := lint.Lint(golangci)
			if len(reports) != test.expectedLen {
				t.Errorf("got %d lint reports, want %d", len(reports), test.expectedLen)
			}
		})
	}
}
