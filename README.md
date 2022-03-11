# Go Vars

This is a Go library for configuration management. It ack as a key-value store
backed with default values. Each value can be override with environment
variable.

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

	// when lookup for environment variable the specified key will transform to
	// snake case in upper case (e.g. MYAPP_KEY1) by default
	vars.SetEnvPrefix("MYAPP_")

	// remove the following code if you don't want to load .env file
	if err := godotenv.Load(); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			log.Fatalln(err)
		}
	}

	// now you can access your configuration with vars.GetXXX("Key")
	// (e.g. vars.GetString("Key1"))
}
```

## License

MIT
