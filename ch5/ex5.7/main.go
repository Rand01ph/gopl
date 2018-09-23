package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		var attrString string
		for _, attr := range n.Attr {
			attrString = fmt.Sprintf(" %s=%s", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attrString)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attrString)
		}
		depth++
	case html.CommentNode:
		fmt.Printf("<!-----%s------>\n", n.Data)
	case html.TextNode:
		txt := strings.TrimSpace(n.Data)
		if len(txt) == 0 {
			return
		}
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	forEachNode(doc, startElement, endElement)
}
