package vars

import (
	"strings"
	"unicode"
)

// Represents a unique key for the value. V is a type of value for this key.
type Key[V any] string

// Provides methods to transform key when lookup on environment variable.
type KeyTransformer interface {
	// Transform the specified key. The return value will be concatenation with
	// prefix without separator.
	Transform(key string) string
}

// A key transformer that convert CamelCase to SNAKE_CASE.
type CamelCaseToSnakeCase struct {
}

func (t *CamelCaseToSnakeCase) Transform(key string) string {
	var builder strings.Builder
	var index int
	var char rune

	start := 0
	write := func() {
		builder.WriteString(strings.ToUpper(key[start:index]))
		if index < len(key) {
			builder.WriteRune('_')
		}
		start = index
	}

	for index, char = range key {
		if unicode.IsUpper(char) {
			if index > 0 {
				write()
			}
		}
	}

	index++
	write()

	return builder.String()
}
