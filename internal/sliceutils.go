package internal

import (
	"cmp"
	"slices"
)

func EqualArray[S ~[]E, E cmp.Ordered](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func IsAlphabetical(original []string) ([]string, bool) {
	sorted := slices.Clone(original)
	slices.Sort(sorted)

	return sorted, EqualArray(original, sorted)
}

func Transform[T any](original []any, f func(a any) (T, bool)) ([]T, bool) {
	ss := make([]T, len(original))
	for i, o := range original {
		s, ok := f(o)
		if !ok {
			return nil, false
		}

		ss[i] = s
	}

	return ss, true
}
