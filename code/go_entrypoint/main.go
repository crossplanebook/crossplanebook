package main

import (
	"fmt"
	"os"
)

var envVars []string

func init() {
	envVars = os.Environ()
}

func main() {
	fmt.Println(envVars)
}
