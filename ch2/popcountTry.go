package main

import (
	"flag"
	"fmt"
	"gopl/ch2/popcount"
	"strconv"
)

var popcountType = flag.String("type", "popcount", "使用的算法")
var popcountNum = flag.Uint64("num", 0, "输入的数")
var p int

func main() {
	flag.Parse()
	for _, num := range flag.Args() {
		t, _ := strconv.ParseUint(num, 10, 64)
		switch *popcountType {
		case "popcount":
			fmt.Printf("popcount %d\n", t)
			p = popcount.PopCount(t)
		case "ex2.3":
			fmt.Printf("ex2.3 %d\n", t)
			p = popcount.PopCount3(t)
		case "ex2.4":
			fmt.Printf("ex2.4 %d\n", t)
			p = popcount.PopCount4(t)
		case "ex2.5":
			fmt.Printf("ex2.5 %d\n", t)
			p = popcount.PopCount5(t)

		default:
			p = popcount.PopCount(t)
		}
		fmt.Printf("pop count is %d\n", p)
	}
}
