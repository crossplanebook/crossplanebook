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
