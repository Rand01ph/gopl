package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("comma is %s\n", comma(arg))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {

	var buf bytes.Buffer
	var before string
	var after string
	var floatString []string

	if strings.HasPrefix(s,"-") || strings.HasPrefix(s,"+"){
		buf.WriteByte(s[0])
		floatString = strings.Split(s[1:], ".")
	} else {
		floatString = strings.Split(s, ".")
	}

	if len(floatString) > 1 {
		before, after = floatString[0], floatString[1]
	} else {
		before, after = floatString[0], ""
	}

	c := len(before) % 3
	fmt.Fprintf(&buf, "%s", before[:c])
	for i := c; i < len(before); i += 3 {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(before[i : i+3])
	}

	if len(after) > 0 {
		buf.WriteByte('.')
		var j int
		for j = 0; j+3 < len(after); j += 3 {
			buf.WriteString(after[j : j+3])
			buf.WriteByte(',')
		}
		buf.WriteString(after[j:])
	}

	return buf.String()
}
