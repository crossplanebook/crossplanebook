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
