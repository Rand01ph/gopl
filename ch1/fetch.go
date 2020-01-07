package main

import (
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport,}
	httpTransport.Dial = dialer.Dial
	for _, url := range os.Args[1:] {
		resp, err := httpClient.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
