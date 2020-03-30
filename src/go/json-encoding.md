# JSON Encoding

JSON encoding may not be commonplace in most introductions to Go (although you
will almost certainly encounter it early on in writing Go programs), but it is
especially important for defining and generating
[CustomResourceDefinitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
in Kubernetes (more on this later). A few important things to keep in mind when
encoding and decoding JSON bytes to structs:

- Only exported fields will be encoded/decoded to JSON.
- Encoding a Go type to JSON is referred to as *marshaling*.
- Decoding JSON bytes to a Go type is referred to as *unmarshaling*.
- [Struct
  tags](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)
  are used to manipulate JSON enconding / decoding (see examples below).
- Unexported fields are initialized to their zero value during unmarshaling.
- An exported field in a Go struct that is not included in the JSON that is
  being decoded are initialized to their zero value during unmarshaling.
- The `omitempty` tag specifies that a field should be omitted from the JSON
  encoding if the field has an empty value, defined as false, 0, a nil pointer,
  a nil interface value, and any empty array, slice, map, or string.
- The `-` tag specifies that a field should be omitted during both encoding and
  decoding.

Let's look at a few examples.

Unexported field is initialized to zero value:
```go
package main

import (
    "encoding/json"
    "fmt"
)

// A Person has qualities of a human
type Person struct {
    First string `json:"first"`
    last  string `json:"last"`
}

func main() {
    p := &Person{}

    j := []byte(`{"first": "Ada", "last": "Lovelace"}`)
    _ = json.Unmarshal(j, p)
    fmt.Println(p.First) // prints "Ada"
    fmt.Println(p.last)  // prints ""
}

```

Exported field with `-` tag is omitted:
```go
package main

import (
    "encoding/json"
    "fmt"
)

// A Person has qualities of a human
type Person struct {
    First string `json:"first"`
    Last  string `json:"-"`
}

func main() {
    p := &Person{}

    j := []byte(`{"first": "Ada", "last": "Lovelace"}`)
    _ = json.Unmarshal(j, p)
    fmt.Println(p.First) // prints "Ada"
    fmt.Println(p.Last)  // prints ""

    p.Last = "Lovelace"

    j, _ = json.Marshal(p)
    fmt.Println(string(j)) // prints {"first":"Ada"}
}

```

Encoding with `omitempty` tag:
```go
package main

import (
    "encoding/json"
    "fmt"
)

// A Person has qualities of a human
type Person struct {
    First string `json:"first"`
    Last  string `json:"last"`
}

// A PersonOmit has qualities of a human
type PersonOmit struct {
    First string `json:"first"`
    Last  string `json:"last,omitempty"`
}

func main() {
    p := &Person{
        First: "Ada",
        Last:  "",
    }

    j, _ := json.Marshal(p)
    fmt.Println(string(j)) // prints {"first":"Ada","last":""}

    o := &PersonOmit{
        First: "Ada",
        Last:  "",
    }

    j, _ = json.Marshal(o)
    fmt.Println(string(j)) // prints {"first":"Ada"}
}
```

For further detail on the JSON package, see the [official
documentation](https://golang.org/pkg/encoding/json/).