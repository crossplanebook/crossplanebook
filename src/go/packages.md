# Packages

Go programs are organized in *packages*.

- Package names must be *lowercase*. You may not include *underscores* or
  *mixedCaps*.
- You may only have one package per directory.
- Packages do not have to match the name of their directory, but it is
  convention to do so.
- Functions and variables that are *uppercase* in a package can be accessed
  directly by a consumer (i.e. they are "public"). These types are referred to
  as *exported*. Functions and variables that are *lowercase* in a package
  cannot be accessed outside of a package (i.e. they are "private"). 
- Exported functions and variables should always have a comment block that
  explains their purposed and function. Most syntax checkers will flag exported
  arguments that do not have a comments block.
- Packages are imported by directory path, but used by package name. For
  instance, if I have a package named `coolpkg` in the
  `github.com/crossplane-book/crossplane-book/pkg/consumed` directory, the
  import will look as follows:

```go

package consumer

import "github.com/crossplane-book/crossplane-book/pkg/consumed" // imported by directory

func myFunc() {
    coolpkg.DoThing() // use by package name [note: public DoThing()]
}
```

- You can alias a package import. This is especially useful when you import two
  packages of the same name:

```go

package consumer

import (
  onetype "github.com/crossplane-book/crossplane-book/pkg/one/type"
  twotype "github.com/crossplane-book/crossplane-book/pkg/two/type"
)

func myFunc() {
    onetype.GetAPIVersion()
    twotype.GetAPIVersion()
}
```

For additional information, checkout the [Go Package Style
Guide](https://rakyll.org/style-packages/).