package vars

import (
	"os"
	"strconv"
)

var defaults = make(map[string]interface{})
var prefix string
var kt KeyTransformer = &CamelCaseToSnakeCase{}
var names = make(map[string]string)
var values = make(map[string]interface{})

// Set default value for the specified key.
func SetDefault(key string, value interface{}) {
	defaults[key] = value
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
	v := getValue(key, func(name, value string) interface{} {
		return value
	})

	return v.(string)
}

// Load the value for the specified and convert it to boolean.
func GetBool(key string) bool {
	v := getValue(key, func(name, value string) interface{} {
		result, err := strconv.ParseBool(value)

		if err != nil {
			panic("Invalid value for environment variable '" + name + "'")
		}

		return result
	})

	return v.(bool)
}

func getValue(key string, parser func(name, value string) interface{}) interface{} {
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
func getEnv(key string, parser func(name, value string) interface{}) (interface{}, bool) {
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
	cache = parser(name, value)
	values[key] = cache

	return cache, true
}
