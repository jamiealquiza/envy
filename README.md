# envy

Simplifies the ability to override Go flags with environment variables. It reduces the need to write conditional checks for each flag and enforces a naming convention.

Envy takes one parameter: a namespace prefix that will be used for environment variable lookups. Each flag registered in your app will be prefixed, uppercased, and hyphens exchanged for underscores; if a matching environment variable is found, it will override (or set) the respective flag value.

### Automatic Env Var Overrides

```go
package main

import (
        "flag"
        "fmt"

        "github.com/jamiealquiza/envy"
)

func main() {
        var address = flag.String("address", "127.0.0.1", "Some random address")
        var port = flag.String("port", "8131", "Some random port")

        envy.Parse("MYAPP") // looks for MYAPP_ADDRESS & MYAPP_PORT
        flag.Parse()

        fmt.Println(*address)
        fmt.Println(*port)
}
```

```
% ./example
127.0.0.1
8131

% MYAPP_ADDRESS="0.0.0.0" MYAPP_PORT="9080" ./example
0.0.0.0
9080
```

### Automatically Adds Env Vars to Help Output

Envy can update your app help output so that it includes the environment variable that would be referenced for overriding each flag. This is done by calling `envy.Parse()` before `flag.Parse()`.

The above example:
```
Usage of ./example:
  -address string
        Some random address [MYAPP_ADDRESS] (default "127.0.0.1")
  -port string
        Some random port [MYAPP_PORT] (default "8131")
```

 If this isn't desired, simply call `envy.Parse()` after `flag.Parse()`:
```go
// ...
		flag.Parse()
        envy.Parse("MYAPP") // looks for MYAPP_ADDRESS & MYAPP_PORT
// ...
```

```
Usage of ./example:
  -address string
        Some random address (default "127.0.0.1")
  -port string
        Some random port (default "8131")
```
