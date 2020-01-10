package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	piper, pipew := io.Pipe()
	go func() {
		defer pipew.Close()
		io.Copy(pipew, reader)
	}()
	io.Copy(os.Stderr, piper)
	piper.Close()
}
