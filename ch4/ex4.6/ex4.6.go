package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "一天学会Go语言        我不信    你信吗"
	b := []byte(s)
	fmt.Println(string(removeSpace(b)))
}

func removeSpace(slice []byte) []byte {
	sum := 0
	for i, b := range slice {
		if unicode.IsSpace(rune(b)) {
			//找到空格移除
			sum ++
		} else {
			if sum > 1 {
				slice = append(slice[:i-sum+1], slice[i:]...)
			}
			sum = 0
		}
	}
	return slice
}
