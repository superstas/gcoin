package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	m := make(map[[sha256.Size]byte]int)
	m[sha256.Sum256([]byte{1})] = 0
	m[sha256.Sum256([]byte{2})] = 0
	m[sha256.Sum256([]byte{3})] = 0
	m[sha256.Sum256([]byte{4})] = 0

	for h := range m {
		fmt.Printf("%x\n", h)
	}
}
