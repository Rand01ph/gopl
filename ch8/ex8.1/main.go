package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"fmt"
	"time"
)

type clock struct {
	tName, tAddr string
}

func main() {
	var tPair []string
	var clockList []clock
	for _, args := range os.Args[1:] {
		fmt.Printf("args is %s\n", args)
		tPair = strings.Split(args, "=")
		fmt.Printf("clock addr is %s\n", tPair[1])
		clockList = append(clockList, clock{tPair[0], tPair[1]})
	}
	for _, c := range clockList {
		conn, err := net.Dial("tcp", c.tAddr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go clockwall(os.Stdout, conn)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func clockwall(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
