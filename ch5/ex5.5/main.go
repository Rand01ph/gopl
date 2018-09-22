package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words num is %d and images num is %d\n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// 图片
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	// 文字
	if n.Type == html.TextNode {
		in := bufio.NewScanner(strings.NewReader(n.Data))
		in.Split(bufio.ScanWords)
		for in.Scan() {
			words++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}
