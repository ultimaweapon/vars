package vars

// Provides methods to convert string to the value.
type ValueParser[T any] interface {
	Parse(s string) (T, error)
}
