package main

import (
	"fmt"
	"unicode/utf8"
)

//先反转rune自身，再反转byte整体

func main() {
	s := "反转Go语言"
	b := []byte(s)
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
	fmt.Println(string(b))

	// 再转回来,遍历然后移动
	for l := len(b); l>0; {
		ru, size := utf8.DecodeRune(b[0:])
		copy(b[0:l], b[0+size:l])
		copy(b[l-size:l], []byte(string(ru)))
		l -= size
	}
	fmt.Println(string(b))
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i ] = s[len(s)-1-i], s[i]
	}
}
