package vars

import (
	"fmt"
	"os"
	"strconv"
)

var defaults = make(map[string]any)
var parsers = make(map[string]any)
var prefix string
var kt KeyTransformer = &CamelCaseToSnakeCase{}
var names = make(map[string]string)
var values = make(map[string]any)

// Set default value for the specified key.
func SetDefault[V any](key Key[V], value V) {
	defaults[string(key)] = value
}

// Set a custom parser to parse the value.
func SetParser[V any](key Key[V], parser ValueParser[V]) {
	parsers[string(key)] = parser
}

// Set prefix for the environment vatiable. The prefix is always treated as
// specified, no automatic transform.
//
// Each key will be transformed when lookup a value from environment variable.
// Use SetEnvKeyTransformer to set what transformer to use. Default is
// CamelCaseToSnakeCase.
func SetEnvPrefix(p string) {
	prefix = p
	names = make(map[string]string)
	values = make(map[string]any)
}

// Set a key transformer to use when lookup on environment variable. Default is
// CamelCaseToSnakeCase.
func SetEnvKeyTransformer(t KeyTransformer) {
	kt = t
	names = make(map[string]string)
	values = make(map[string]any)
}

// Load the value for the specified key.
func Get[V any](key Key[V]) V {
	// environment variable
	value, exists := getEnv(key)

	if exists {
		return value
	}

	// default
	def, exists := defaults[string(key)]

	if exists {
		return def.(V)
	}

	panic("Key '" + key + "' does not exists")
}

// We don't care for race condition here due to all mutations come from pure
// functions.
func getEnv[V any](key Key[V]) (V, bool) {
	// lookup cache
	cache, exists := values[string(key)]

	if exists {
		return cache.(V), true
	}

	// get name
	name, exists := names[string(key)]

	if !exists {
		name = prefix + kt.Transform(string(key))
		names[string(key)] = name
	}

	// load variable
	value, exists := os.LookupEnv(name)

	if !exists {
		// FIXME: remove zero without introduce named result parameter
		var zero V
		return zero, false
	}

	// parse variable
	var err error

	parser, exists := parsers[string(key)]

	if exists {
		cache, err = parser.(ValueParser[V]).Parse(value)
	} else {
		// FIXME: use V instead
		switch defaults[string(key)].(type) {
		case string:
			cache = value
		case bool:
			cache, err = strconv.ParseBool(value)
		case int:
			var res int64

			if res, err = strconv.ParseInt(value, 0, strconv.IntSize); err == nil {
				cache = int(res)
			}
		default:
			panic("No parser for '" + key + "'")
		}
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to parse environment variable '%v': %v", name, err))
	}

	values[string(key)] = cache

	return cache.(V), true
}
