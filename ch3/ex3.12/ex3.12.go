package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	a := strings.Split(os.Args[1], "")
	sort.Strings(a)
	first := strings.Join(a, "")

	b := strings.Split(os.Args[2], "")
	sort.Strings(b)
	second := strings.Join(b, "")

	if len(first) != len(second) {
		fmt.Printf("not equal\n")
	} else {
		for i := 0; i < len(first); i++ {
			if first[i] != second[i] {
				fmt.Printf("not equal\n")
				os.Exit(1)
			}
		}
		fmt.Printf("equal\n")
	}

}
