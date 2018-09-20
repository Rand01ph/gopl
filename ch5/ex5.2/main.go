package main

import (
	"os"
	"fmt"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.2: %v\n", err)
		os.Exit(1)
	}
	elemMap := make(map[string]int)

	for elem, sum := range countElem(elemMap, doc) {
		fmt.Printf("%s\t%d\n", elem, sum)
	}
}

func countElem(eMap map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		eMap[n.Data]++
	}

	if n.NextSibling != nil {
		eMap = countElem(eMap, n.NextSibling)
	}
	if n.FirstChild != nil {
		eMap = countElem(eMap, n.FirstChild)
	}
	return eMap

}

