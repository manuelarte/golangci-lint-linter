package internal

import "github.com/goccy/go-yaml"

type (
	Golangci struct {
		fields yaml.MapSlice

		cm map[string][]*yaml.Comment
	}

	Linters struct {
		fields yaml.MapSlice

		cm map[string][]*yaml.Comment
	}
)

func (g Golangci) GetLinters() (Linters, bool) {
	linters, ok := getKey[yaml.MapSlice](g.fields, "linters")
	if !ok {
		return Linters{}, false
	}

	return Linters{
		fields: linters,
		cm:     g.cm,
	}, true
}

func (l Linters) GetEnable() ([]string, bool) {
	enableInterface, ok := getKey[[]any](l.fields, "enable")
	if !ok {
		return nil, false
	}

	return Transform[string](enableInterface, anyToString)
}

func (l Linters) GetDisable() ([]string, bool) {
	disableInterface, ok := getKey[[]any](l.fields, "disable")
	if !ok {
		return nil, false
	}

	return Transform[string](disableInterface, anyToString)
}

func Parse(input []byte) (Golangci, error) {
	cm := make(map[string][]*yaml.Comment)

	var fields yaml.MapSlice

	err := yaml.UnmarshalWithOptions(input, &fields, yaml.UseOrderedMap(), yaml.CommentToMap(cm))
	if err != nil {
		return Golangci{}, err
	}

	return Golangci{
		fields: fields,
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

		value, isT := field.Value.(T)
		if !isT {
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
