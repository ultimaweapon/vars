package vars

import "testing"

func TestCamelCaseToSnakeCase(t *testing.T) {
	kt := CamelCaseToSnakeCase{}

	t.Run("Foo", func(t *testing.T) {
		if r := kt.Transform("Foo"); r != "FOO" {
			t.Errorf("Transform(Foo) = %v; want FOO", r)
		}
	})

	t.Run("FooBar", func(t *testing.T) {
		if r := kt.Transform("FooBar"); r != "FOO_BAR" {
			t.Errorf("Transform(FooBar) = %v; want FOO_BAR", r)
		}
	})
}
