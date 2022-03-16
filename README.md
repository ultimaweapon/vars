# Go Vars
[![Go Reference](https://pkg.go.dev/badge/github.com/ultimicro/vars.svg)](https://pkg.go.dev/github.com/ultimicro/vars)

This is a Go library for simple configuration management. It act as a key-value store backed with default value. Each
value can be override with environment variable. This library utilize generics in Go 1.18.

## Usage

```go
package main

import (
	"errors"
	"io/fs"
	"log"

	"github.com/joho/godotenv"
	"github.com/ultimicro/vars"
)

const (
	KeyFoo vars.Key[int]    = "Foo"
	KeyBar vars.Key[string] = "Bar"
)

func main() {
	vars.SetDefault(KeyFoo, 7)
	vars.SetDefault(KeyBar, "abc")

	// When lookup for environment variable the specified key will transform to
	// snake case in upper case (e.g. MYAPP_KEY1) by default. Use
	// SetEnvKeyTransformer() to change this behavior.
	vars.SetEnvPrefix("MYAPP_")

	// Remove the following code if you don't want to load .env file.
	if err := godotenv.Load(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			log.Fatalln(err)
		}
	}

	// Now you can access your configuration with vars.Get(Key).
}
```

### Custom parser

If you need to change how the value is parsed you can set an implementation of `ValueParser` to use with `SetParser`:

```go
vars.SetParser[SomeType](Key, &SomeTypeParser{})
```

Yes you need to specify type manually due to currently Go cannot infer type in this case.

## License

MIT
