package vars

import (
	"fmt"
	"os"
	"strconv"
)

var defaults = make(map[string]interface{})
var parsers = make(map[string]ValueParser)
var prefix string
var kt KeyTransformer = &CamelCaseToSnakeCase{}
var names = make(map[string]string)
var values = make(map[string]interface{})

// Set default value for the specified key. Options can be a ValueParser.
func SetDefault(key string, value interface{}, options ...interface{}) {
	defaults[key] = value

	for _, option := range options {
		switch v := option.(type) {
		case ValueParser:
			parsers[key] = v
		default:
			panic(fmt.Sprintf("Unknown option '%T'", v))
		}
	}
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
	values = make(map[string]interface{})
}

// Set a key transformer to use when lookup on environment variable. Default is
// CamelCaseToSnakeCase.
func SetEnvKeyTransformer(t KeyTransformer) {
	kt = t
	names = make(map[string]string)
	values = make(map[string]interface{})
}

// Load the value for the specified key and treated it as a string.
func GetString(key string) string {
	v := getValue(key, func(s string) (interface{}, error) {
		return s, nil
	})

	return v.(string)
}

// Load the value for the specified key and convert it to bool.
func GetBool(key string) bool {
	v := getValue(key, func(s string) (interface{}, error) {
		return strconv.ParseBool(s)
	})

	return v.(bool)
}

// Load the value for the specified key and convert it to int.
func GetInt(key string) int {
	v := getValue(key, func(s string) (interface{}, error) {
		if result, err := strconv.ParseInt(s, 0, strconv.IntSize); err != nil {
			return nil, err
		} else {
			return int(result), nil
		}
	})

	return v.(int)
}

func getValue(key string, parser func(s string) (interface{}, error)) interface{} {
	// environment variable
	value, exists := getEnv(key, parser)

	if exists {
		return value
	}

	// default
	value, exists = defaults[key]

	if exists {
		return value
	}

	panic("Key '" + key + "' does not exists")
}

// We don't care for race condition here due to all mutations come from pure
// functions.
func getEnv(key string, fallback func(s string) (interface{}, error)) (interface{}, bool) {
	// lookup cache
	cache, exists := values[key]

	if exists {
		return cache, true
	}

	// get name
	name, exists := names[key]

	if !exists {
		name = prefix + kt.Transform(key)
		names[key] = name
	}

	// load variable
	value, exists := os.LookupEnv(name)

	if !exists {
		return nil, false
	}

	// parse variable
	var err error

	parser, exists := parsers[key]

	if exists {
		cache, err = parser.Parse(value)
	} else {
		cache, err = fallback(value)
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to parse environment variable '%v': %v", name, err))
	}

	values[key] = cache

	return cache, true
}
