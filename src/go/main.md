# An Entrypoint

Like most programming languages, Go requires an entrypoint to get a program
started. You may have multiple entrypoints in a single project, but only one per
package and the package name must be `main`. Inside of a `main` package, the
`main()` function is used to start the program. However, in a non-`main`
package, `main()` is not reserved as a special word and can be used as a normal
function or variable.


```go
package main

func main() {
    // do things here
}
```

Another special function to the `main` package is `init()`. The `init()`
executes prior to `main()`, so it is commonly used for operations like reading
and setting environment variables, parsing command flags, and opening database
connections.

```go
package main

import (
    "fmt"
    "os"
)

var envVars []string

func init() { // this is run first
    envVars = os.Environ()
}

func main() { // evnVars is populated when main() runs
    fmt.Println(envVars)
}
```

Note that neither the `main()` nor the `init()` function ever have explicit
return values. If you need to exit `main()` early, you can use
[`os.Exit()`](https://gobyexample.com/exit).