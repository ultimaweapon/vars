package vars

import "testing"

func TestCamelCaseToSnakeCase(t *testing.T) {
	kt := CamelCaseToSnakeCase{}
	cases := [...]struct {
		Input  string
		Expect string
	}{
		{
			Input:  "Foo",
			Expect: "FOO",
		},
		{
			Input:  "FooBar",
			Expect: "FOO_BAR",
		},
	}

	for _, c := range cases {
		r := kt.Transform(c.Input)

		if r != c.Expect {
			t.Errorf("Transform(%v) = %v; want %v", c.Input, r, c.Expect)
		}
	}
}
