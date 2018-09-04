package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	sum := 0
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			sum += 1
		}
	}
	fmt.Printf("different sum is %d", sum)
}