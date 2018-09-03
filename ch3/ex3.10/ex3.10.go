package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("comma is %s\n", comma(arg))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	c := len(s) % 3
	fmt.Fprintf(&buf, "%s", s[:c])
	for i := c; i < len(s); i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
