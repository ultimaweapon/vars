# Go Vars
[![Go Reference](https://pkg.go.dev/badge/github.com/ultimicro/vars.svg)](https://pkg.go.dev/github.com/ultimicro/vars)

This is a Go library for configuration management. It act as a key-value store backed with default value. Each value can
be override with environment variable.

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

func main() {
	vars.SetDefault("Key1", "Foo")
	vars.SetDefault("Key2", false)

	// When lookup for environment variable the specified key will transform to snake case in upper case (e.g. MYAPP_KEY1)
	// by default. Use SetEnvKeyTransformer() to change this behavior.
	vars.SetEnvPrefix("MYAPP_")

	// Remove the following code if you don't want to load .env file.
	if err := godotenv.Load(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			log.Fatalln(err)
		}
	}

	// Now you can access your configuration with vars.GetXXX("Key") (e.g. vars.GetString("Key1")).
}
```

### Custom parser

If you need to change how the value is parsed you can provide an implementation of `ValueParser` next to the default
value:

```go
vars.SetDefault("Key", SomeConstant, &ParseSomeConstant)
```

## License

MIT
