package vars

import (
	"os"
	"strings"
	"testing"
)

func TestDefault(t *testing.T) {
	var key Key[bool] = "Foo"

	t.Cleanup(clearStates)

	SetDefault(key, true)

	if r := Get(key); r != true {
		t.Errorf("Get(Foo) = %v; want true", r)
	}
}

func TestCustomParser(t *testing.T) {
	var key Key[string] = "StringKey"

	t.Cleanup(clearStates)
	t.Cleanup(func() {
		os.Unsetenv("STRING_KEY")
	})

	os.Setenv("STRING_KEY", "abc")

	SetDefault(key, "")
	SetParser[string](key, &stringParser{})

	if r := Get(key); r != "ABC" {
		t.Errorf("Get(StringKey) = %v; want ABC", r)
	}
}

func TestEnvPrefix(t *testing.T) {
	var key Key[int] = "Foo"

	t.Cleanup(clearStates)
	t.Cleanup(func() {
		os.Unsetenv("MYAPP_FOO")
	})

	os.Setenv("MYAPP_FOO", "9")

	SetDefault(key, 7)
	SetEnvPrefix("MYAPP_")

	if r := Get(key); r != 9 {
		t.Errorf("Get(Foo) = %v; want 9", r)
	}
}

func clearStates() {
	defaults = make(map[string]any)
	parsers = make(map[string]any)
	prefix = ""
	kt = &CamelCaseToSnakeCase{}
	names = make(map[string]string)
	values = make(map[string]any)
}

type stringParser struct {
}

func (p *stringParser) Parse(s string) (string, error) {
	return strings.ToUpper(s), nil
}
