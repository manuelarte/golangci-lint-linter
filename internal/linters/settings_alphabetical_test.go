package linters

import (
	"testing"

	"github.com/manuelarte/golangci-lint-linter-go/internal"
	"reflect"
)

func TestSettingsAlphabetical_Lint(t *testing.T) {
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
		"linters settings not sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  settings:
    bidichk:
      left-to-right-embedding: false
    asasalint:
      use-builtin-exclusions: false
`),
			expectedLen: 1,
		},
		"linters settings sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  settings:
    asasalint:
      use-builtin-exclusions: false
    bidichk:
      left-to-right-embedding: false
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

			lint := NewSettingsAlphabetical()

			reports := lint.Lint(golangci)
			if len(reports) != test.expectedLen {
				t.Errorf("got %d lint reports, want %d", len(reports), test.expectedLen)
			}
		})
	}
}

func TestSettingsAlphabetical_Fix(t *testing.T) {
	testCases := map[string]struct {
		input    []byte
		expected []byte
	}{
		"linters field not present": {
			input: []byte(`
version: 2
`),
			expected: []byte("version: 2\n"),
		},
		"linters.settings field not present": {
			input: []byte(`
version: 2
linters:
  enable:
    - funcorder
`),
			expected: []byte(`version: 2
linters:
  enable:
    - funcorder
`),
		},
		"linters settings not sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  enable:
    - funcorder    
    - tagliatelle # make sure all the json tags are following the same pattern
  settings:
    tagliatelle:
      propTagliatelle: valueTagliatelle
    funcorder:
      propFuncorder: valueFuncorder
`),
			expected: []byte(
				`version: 2
linters:
  enable:
    - funcorder
    - tagliatelle # make sure all the json tags are following the same pattern
  settings:
    funcorder:
      propFuncorder: valueFuncorder
    tagliatelle:
      propTagliatelle: valueTagliatelle
`),
		},
		"linters settings sorted alphabetically": {
			input: []byte(`
version: 2
linters:
  enable:
    - funcorder    
    - tagliatelle # make sure all the json tags are following the same pattern
  settings:
    funcorder:
      propFuncorder: valueFuncorder    
    tagliatelle:
      propTagliatelle: valueTagliatelle
`),
			expected: []byte(
				`version: 2
linters:
  enable:
    - funcorder
    - tagliatelle # make sure all the json tags are following the same pattern
  settings:
    funcorder:
      propFuncorder: valueFuncorder
    tagliatelle:
      propTagliatelle: valueTagliatelle
`),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			golangci, err := internal.Parse(test.input)
			if err != nil {
				t.Fatal(err)
			}

			lint := NewSettingsAlphabetical()
			err = lint.Fix(golangci)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := golangci.(internal.YamlGolangci).Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("got %#v, want %#v", string(actual), string(test.expected))
			}
		})
	}
}
