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
