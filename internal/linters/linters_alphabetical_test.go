package linters

import (
	"testing"

	"github.com/manuelarte/golangci-lint-linter-go/internal"
)

func TestLintersAlphabetical_Lint(t *testing.T) {
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
		"linters enable not sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  enable:
    - tagliatelle
    - funcorder
`),
			expectedLen: 1,
		},
		"linters enable sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  enable:
    - funcorder
    - tagliatelle
`),
			expectedLen: 0,
		},
		"linters disable not sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  disable:
    - tagliatelle
    - funcorder
`),
			expectedLen: 1,
		},
		"linters disable sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  disable:
    - funcorder
    - tagliatelle
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

			lint := NewLintersAlphabetical()

			reports := lint.Lint(golangci)
			if len(reports) != test.expectedLen {
				t.Errorf("got %d lint reports, want %d", len(reports), test.expectedLen)
			}
		})
	}
}
