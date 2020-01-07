// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2020-01-07 21:48:46

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Provide custom arguments to compile with Alt-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
	"path"
	"io"
	"os"
	"net/http"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	f, n, err := fetch("https://www.baidu.com/")
	if err != nil {
		fmt.Printf("the fetch err is %v\n", err)
	} else {
		fmt.Printf("the fetch file is %s and size is %d\n", f, n)
	}
}
