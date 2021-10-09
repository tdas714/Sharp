package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// Print Hello world
	sum := "Hello World"
	fmt.Printf("%x", sha256.Sum256([]byte(sum)))
}
