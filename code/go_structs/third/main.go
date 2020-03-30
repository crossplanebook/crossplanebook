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
