package internal

import (
	"github.com/goccy/go-yaml"
)

var (
	_ Golangci = new(YamlGolangci)
	_ Linters  = new(YamlLinters)
	_ Settings = new(YamlSettings)
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
		GetSettings() (Settings, bool)
	}

	Settings interface {
		GetSetting(key string) (interface{}, bool)
		GetKeys() ([]string, bool)
	}

	YamlGolangci struct {
		fields *yaml.MapSlice

		cm map[string][]*yaml.Comment
	}

	YamlLinters struct {
		fields *yaml.MapSlice

		cm map[string][]*yaml.Comment
	}

	YamlSettings struct {
		fields *yaml.MapSlice

		cm map[string][]*yaml.Comment
	}
)

func (g YamlGolangci) GetLinters() (Linters, bool) {
	linters, ok := getKey[yaml.MapSlice](*g.fields, "linters")
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
	return yaml.MarshalWithOptions(g.fields, yaml.WithComment(g.cm))
}

func (l YamlLinters) GetEnable() ([]string, bool) {
	enableInterface, ok := getKey[[]any](*l.fields, "enable")
	if !ok {
		return nil, false
	}

	return Transform[string](enableInterface, anyToString)
}

func (l YamlLinters) GetDisable() ([]string, bool) {
	disableInterface, ok := getKey[[]any](*l.fields, "disable")
	if !ok {
		return nil, false
	}

	return Transform[string](disableInterface, anyToString)
}

func (l YamlLinters) GetSettings() (Settings, bool) {
	settings, ok := getKey[yaml.MapSlice](*l.fields, "settings")
	if !ok {
		return nil, false
	}

	return YamlSettings{
		fields: &settings,
		cm:     l.cm,
	}, true
}

func (y YamlSettings) GetSetting(key string) (any, bool) {
	return getKey[map[string]any](*y.fields, key)
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

func getKey[T any](fields yaml.MapSlice, key string) (T, bool) {
	var zero T

	for _, field := range fields {
		k, isString := field.Key.(string)
		if !isString {
			return zero, false
		}

		if key != k {
			continue
		}

		value, isTypeT := field.Value.(T)
		if !isTypeT {
			return zero, false
		}

		return value, true
	}

	return zero, false
}

func anyToString(a any) (string, bool) {
	casted, isString := a.(string)

	return casted, isString
}
