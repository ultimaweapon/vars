package vars

// Provides methods to convert string to the value.
type ValueParser interface {
	Parse(s string) (interface{}, error)
}
