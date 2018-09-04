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

	if first == second {
		fmt.Printf("equal\n")
	} else {
		fmt.Printf("not equal\n")
	}

}
