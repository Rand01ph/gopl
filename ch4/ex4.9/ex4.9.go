package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		wordfreq[in.Text()]++
	}
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("wordfreq\tcount\n")
	for c, n := range wordfreq {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
