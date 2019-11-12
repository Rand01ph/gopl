package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/proxy"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if a := forEachNode(c, id, pre, post); a != nil {
			return a
		}
	}
	if post != nil {
		if post(n, id) {
			return n
		}
	}
	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
	}
	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

func main() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	resp, err := httpClient.Get(os.Args[1])
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	wanted := ElementByID(doc, os.Args[2])
	fmt.Printf("the wanted node is %v", wanted)
}
