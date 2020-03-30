package main

import "fmt"

// A Being can return its first and last name
type Being interface {
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

// A Dog is a type of animal
type Dog struct {
	name  string
	breed string
}

// GetFirstName get the first name of a Dog
func (d *Dog) GetFirstName() string {
	return d.name
}

// GetLastName get the first name of a Dog
func (d *Dog) GetLastName() string {
	// structs may implement an interface however they see fit
	return ""
}

// GetFullName returns the full name for a Being
func GetFullName(b Being) string {
	// Passing a Being means we can contruct the fullname for any type that
	// implements GetFirstName and GetLastName
	return fmt.Sprintf("%s %s", b.GetFirstName(), b.GetLastName())
}

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
	}

	fmt.Println(GetFullName(p)) // prints "Ada Lovelace"
	fmt.Println(GetFullName(d)) // prints "Spot "
}
