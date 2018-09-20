package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
	"flag"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.4: %v\n", err)
		os.Exit(1)
	}
	elemType := flag.String("t", "a", "输入想要查询的结点类型")
	flag.Parse()
	for _, link := range visit(nil, doc, *elemType) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node, elemType string) []string {
	if n.Type == html.ElementNode && n.Data == elemType {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling, elemType)
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild, elemType)
	}
	return links
}
