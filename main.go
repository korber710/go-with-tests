package main

import (
	"fmt"
	dependency_injection "github.com/korber710/go-with-tests/pkg/dependency-injection"
	"os"
)

func main() {
	fmt.Println("This is for practicing Go with tests")
	dependency_injection.Greet(os.Stdout, "Steve")
}
