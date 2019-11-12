package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	sSlice := strings.Split(s, " ")
	for i, _ := range sSlice {
		sSlice[i] = f(sSlice[i])
	}
	return strings.Join(sSlice, " ")
}

func printS(s string) string {
	return fmt.Sprintf("***%s***", s)
}

func main() {
	a := "aaa bbb ccc"
	fmt.Println(expand(a, printS))
}