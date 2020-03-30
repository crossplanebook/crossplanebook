package main

import "fmt"

// Aged is something that has an age
type Aged interface {
	GetAge() int
}

// A Being can return its first and last name
type Being interface {
	Aged
	GetFirstName() string
	GetLastName() string
}

// An Person is a human
type Person struct {
	first       string
	last        string
	age         int
	nationality string
}

// GetFirstName get the first name of a Person
func (p *Person) GetFirstName() string {
	return p.first
}

// GetLastName get the first name of a Person
func (p *Person) GetLastName() string {
	return p.last
}

// GetAge gets the age of a Person
func (p *Person) GetAge() int {
	return p.age
}

// A Dog is a type of animal
type Dog struct {
	name  string
	breed string
	age   int
}

// GetFirstName get the first name of a Dog
func (d *Dog) GetFirstName() string {
	return d.name
}

// GetLastName get the first name of a Dog
func (d *Dog) GetLastName() string {
	return ""
}

// GetAge gets the age of a Dog
func (d *Dog) GetAge() int {
	return d.age
}

// GetFullName returns the full name for a Being
func GetFullName(b Being) string {
	return fmt.Sprintf("%s %s", b.GetFirstName(), b.GetLastName())
}

// GetFullNameAndAge returns the full name and age for a Being
func GetFullNameAndAge(b Being) string {
	return fmt.Sprintf("%s %s: %d", b.GetFirstName(), b.GetLastName(), b.GetAge())
}

// blank identifiers are useful for ensuring a type satisfies an interface at
// compile time
var _ Being = &Person{}
var _ Being = &Dog{}

func main() {
	p := &Person{
		first:       "Ada",
		last:        "Lovelace",
		age:         36,
		nationality: "British",
	}

	d := &Dog{
		name:  "Spot",
		breed: "Corgi",
		age:   2,
	}

	fmt.Println(GetFullNameAndAge(p)) // prints "Ada Lovelace: 36"
	fmt.Println(GetFullNameAndAge(d)) // prints "Spot : 2"
}
