package main

import "fmt"

func clear(slice []string) []string {
	for i:=0; i < len(slice)-1; i++ {
		k := 0
		for j:=i+1; slice[j] == slice[i] ; {
			k++
			j++
			if j == len(slice) {
				break
			}
		}
		if k > 0 {
			slice = append(slice[:i], slice[i+k:]...)
		}
	}
	return slice
}

func main() {
	s := []string{"a", "a", "a", "a", "b", "b", "c", "d", "d", "e", "e", "e"}
	fmt.Println(clear(s[:]))
}
