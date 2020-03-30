# Structs

If you have a background in Object Oriented Program, you are likely familiar
with the concept of classes. Go does not provide classes, but instead defines
typed collections of fields as *structs*. A struct can have both field and
methods. Structs can also embed other structs.

Let's look at an example of a struct with fields:

```go
package main

import "fmt"

// A Person has qualities of a human
type Person struct {
    first string
    last  string
}

func main() {
    p := &Person{ // you can create new struct instances with :=
        first: "Ada",
        last:  "Lovelace",
    }

    fmt.Println(p.first) // we can access "first" because Person is defined in this package
}
```

Now let's take a look at a struct with a method:

```go
package main

import "fmt"

// A Person has qualities of a human
type Person struct {
	first string
	last  string
}

// GetFirstName returns the first name of a person
func (p *Person) GetFirstName() string { // methods are like functions but have a receiver: (p *Person)
	return p.first
}

func main() {
	p := &Person{
		first: "Ada",
		last:  "Lovelace",
	}

	fmt.Println(p.GetFirstName()) // We could still get the Person first name outside of this package because GetFirstName is public
}
```

Lastly, let's see what embedding a struct in another struct looks like:

```go
package main

import "fmt"

// A Person has qualities of a human
type Person struct {
	first string
	last  string
}

// GetFirstName returns the first name of a person
func (p *Person) GetFirstName() string {
	return p.first
}

// An Athlete is a superset of a Person
type Athlete struct {
	Person
	sport string
}

func main() {
	p := &Athlete{
		Person: Person{
			first: "LeBron",
			last:  "James",
		},
		sport: "basketball",
	}

	fmt.Println(p.GetFirstName())        // We do not have to access the Person sub-struct to call its methods
	fmt.Println(p.Person.GetFirstName()) // We can explicitly reference the Person sub-struct and its methods

	// Note: if Athlete had its own GetFirstName() method, then calling
	// p.GetFirstName() would invoke it and p.Person.GetFirstName() would invoke
	// the Person implementation.
}
```