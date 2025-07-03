package internal

import (
	"bytes"
	"testing"
)

func TestRouncTrip(t *testing.T) {
	testCases := map[string]struct {
		input    []byte
		expected []byte
	}{
		"simple file no comments": {
			input: []byte(`
version: 2
`),
			expected: []byte("version: 2\n"),
		},
		"simple file with comments": {
			input: []byte(`
version: 2 # the version
`),
			expected: []byte("version: 2 # the version\n"),
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			golangci, err := Parse(test.input)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := golangci.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(actual, test.expected) {
				t.Errorf("actual = %q, expected = %q", actual, test.expected)
			}
		})
	}
}
