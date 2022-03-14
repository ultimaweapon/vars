package vars

import (
	"os"
	"strings"
	"testing"
)

func TestValueParser(t *testing.T) {
	t.Cleanup(clearStates)
	t.Cleanup(func() {
		os.Unsetenv("STRING_KEY")
	})

	os.Setenv("STRING_KEY", "abc")

	SetDefault("StringKey", "", &stringParser{})

	if r := GetString("StringKey"); r != "ABC" {
		t.Errorf("GetString(StringKey) = %v; want ABC", r)
	}
}

func clearStates() {
	defaults = make(map[string]interface{})
	parsers = make(map[string]ValueParser)
	prefix = ""
	kt = &CamelCaseToSnakeCase{}
	names = make(map[string]string)
	values = make(map[string]interface{})
}

type stringParser struct {
}

func (p *stringParser) Parse(s string) (interface{}, error) {
	return strings.ToUpper(s), nil
}
