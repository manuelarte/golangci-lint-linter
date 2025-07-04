package internal

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"cmp"
	"github.com/goccy/go-yaml"
)

var errLintersEnableNotFound = errors.New("linters.enable not found")

var (
	_ Golangci = new(YamlGolangci)
	_ Linters  = new(YamlLinters)
	_ Settings = new(YamlSettings)

	//nolint:gochecknoglobals // ignore this
	cmEnablePath = func(index int) string { return fmt.Sprintf("$.linters.enable[%d]", index) }
	//nolint:gochecknoglobals // ignore this
	cmDisablePath = func(index int) string { return fmt.Sprintf("$.linters.disable[%d]", index) }
)

type (
	Golangci interface {
		GetLinters() (Linters, bool)
		GetComment(path string) (string, bool)
		Marshal() ([]byte, error)
	}

	Linters interface {
		GetEnable() ([]string, bool)
		GetDisable() ([]string, bool)
		GetFieldNames() ([]string, bool)
		GetSettings() (Settings, bool)

		SortEnable() error
		SortDisable() error
		SortFields(order map[string]int) error
		SortSettings() error
	}

	Settings interface {
		GetSetting(key string) (any, bool)
		GetKeys() ([]string, bool)
		Sort() error
	}

	YamlGolangci struct {
		fields *yaml.MapSlice
		cm     map[string][]*yaml.Comment
	}

	YamlLinters struct {
		fields *yaml.MapSlice
		cm     map[string][]*yaml.Comment
	}

	YamlSettings struct {
		fields *yaml.MapSlice
		cm     map[string][]*yaml.Comment
	}
)

func (g YamlGolangci) GetLinters() (Linters, bool) {
	linters, _, ok := getKey[yaml.MapSlice](*g.fields, "linters")
	if !ok {
		return YamlLinters{}, false
	}

	return YamlLinters{
		fields: &linters,
		cm:     g.cm,
	}, true
}

func (g YamlGolangci) GetComment(path string) (string, bool) {
	comments, ok := g.cm[path]
	if !ok {
		return "", false
	}

	for _, comment := range comments {
		if comment.Position == yaml.CommentLinePosition {
			return comment.Texts[0], true
		}
	}

	return "", false
}

func (g YamlGolangci) Marshal() ([]byte, error) {
	return yaml.MarshalWithOptions(g.fields, yaml.IndentSequence(true), yaml.WithComment(g.cm))
}

func (l YamlLinters) GetEnable() ([]string, bool) {
	enableInterface, _, ok := getKey[[]any](*l.fields, "enable")
	if !ok {
		return nil, false
	}

	return Transform[string](enableInterface, anyToString)
}

func (l YamlLinters) GetDisable() ([]string, bool) {
	disableInterface, _, ok := getKey[[]any](*l.fields, "disable")
	if !ok {
		return nil, false
	}

	return Transform[string](disableInterface, anyToString)
}

func (l YamlLinters) GetFieldNames() ([]string, bool) {
	fields := make([]string, 0, len(*l.fields))
	for _, field := range *l.fields {
		if s, isString := field.Key.(string); isString {
			fields = append(fields, s)
		}
	}
	return fields, true
}

func (l YamlLinters) GetSettings() (Settings, bool) {
	settings, _, ok := getKey[yaml.MapSlice](*l.fields, "settings")
	if !ok {
		return nil, false
	}

	return YamlSettings{
		fields: &settings,
		cm:     l.cm,
	}, true
}

func (l YamlLinters) SortEnable() error {
	enable, ok := l.GetEnable()
	if !ok {
		return nil
	}

	sorted, hasSorted := l.sort(enable, cmEnablePath)
	if !hasSorted {
		return nil
	}

	return l.replace("enable", sorted)
}

func (l YamlLinters) SortDisable() error {
	disable, ok := l.GetDisable()
	if !ok {
		return nil
	}

	sorted, hasSorted := l.sort(disable, cmDisablePath)
	if !hasSorted {
		return nil
	}

	return l.replace("disable", sorted)
}

func (l YamlLinters) SortFields(expectedOrder map[string]int) error {
	slices.SortFunc(*l.fields, func(a, b yaml.MapItem) int {

		aString, isAString := a.Key.(string)
		bString, isBString := b.Key.(string)

		if isAString && isBString {
			aIndex, hasA := expectedOrder[a.Key.(string)]
			bIndex, hasB := expectedOrder[b.Key.(string)]
			if !hasA && !hasB {
				return cmp.Compare(aIndex, bIndex)
			} else {
				strings.Compare(aString, bString)
			}
		}

		return 0
	})
	return nil
}

func (l YamlLinters) SortSettings() error {
	settings, ok := l.GetSettings()
	if !ok {
		return nil
	}

	return settings.Sort()
}

func (l YamlLinters) sort(original []string, cmPath func(index int) string) ([]string, bool) {
	sorted, isSorted := IsAlphabetical(original)
	if isSorted {
		return nil, false
	}

	commentIndexedByLinter := l.indexLintersComments(original, cmPath)

	for sortedIndex, linter := range sorted {
		originalIndex := slices.Index(original, linter)
		if comment, hasComment := commentIndexedByLinter[linter]; hasComment && originalIndex != sortedIndex {
			l.cm[cmPath(sortedIndex)] = comment
			delete(l.cm, cmPath(originalIndex))
		}
	}

	return sorted, true
}

func (l YamlLinters) indexLintersComments(original []string, cmPath func(index int) string) map[string][]*yaml.Comment {
	indexMap := make(map[string][]*yaml.Comment)

	for index, linter := range original {
		if comment, hasComment := l.cm[cmPath(index)]; hasComment {
			indexMap[linter] = comment
		}
	}

	return indexMap
}

func (l YamlLinters) replace(path string, n []string) error {
	_, index, ok := getKey[[]any](*l.fields, path)
	if !ok {
		return errLintersEnableNotFound
	}

	v := *l.fields
	original := v[index]
	v[index] = yaml.MapItem{
		Key:   original.Key,
		Value: n,
	}

	return nil
}

func (y YamlSettings) GetSetting(key string) (any, bool) {
	r, _, ok := getKey[map[string]any](*y.fields, key)

	return r, ok
}

func (y YamlSettings) GetKeys() ([]string, bool) {
	keys := make([]string, 0, len(*y.fields))
	for _, k := range *y.fields {
		if casted, ok := k.Key.(string); ok {
			keys = append(keys, casted)
		}
	}

	return keys, true
}

func (y YamlSettings) Sort() error {
	_, ok := y.GetKeys()
	if !ok {
		return nil
	}

	slices.SortFunc(*y.fields, func(a, b yaml.MapItem) int {
		//nolint:errcheck // already checked with GetKeys
		return strings.Compare(a.Key.(string), b.Key.(string))
	})

	return nil
}

func Parse(input []byte) (Golangci, error) {
	cm := make(map[string][]*yaml.Comment)

	var fields yaml.MapSlice

	err := yaml.UnmarshalWithOptions(input, &fields, yaml.UseOrderedMap(), yaml.CommentToMap(cm))
	if err != nil {
		return YamlGolangci{}, err
	}

	return YamlGolangci{
		fields: &fields,
		cm:     cm,
	}, nil
}

func getKey[T any](fields yaml.MapSlice, key string) (T, int, bool) {
	var zero T

	for i, field := range fields {
		k, isString := field.Key.(string)
		if !isString {
			return zero, 0, false
		}

		if key != k {
			continue
		}

		value, isTypeT := field.Value.(T)
		if !isTypeT {
			return zero, 0, false
		}

		return value, i, true
	}

	return zero, 0, false
}

func anyToString(a any) (string, bool) {
	casted, isString := a.(string)

	return casted, isString
}
