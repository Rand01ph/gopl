package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			f_ret := strings.Split(line, ":")
			file_name, file_content := f_ret[0], f_ret[1]
			if file_name != "/dev/stdin" {
				fmt.Printf("%d\t%s\t%s\n", n, file_content, file_name)
			} else {
				fmt.Printf("%d\t%s\n", n, file_content)
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[f.Name()+":"+input.Text()]++
	}
}
