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
