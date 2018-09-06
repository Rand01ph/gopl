package main

import "fmt"

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	b := rotate(a[:], 2)
	fmt.Println(b)
}

func rotate(slice []int, num int) []int {
	for i := 0; i < num; i++ {
		slice = append(slice, slice[i])
	}
	return slice[num:]
}
