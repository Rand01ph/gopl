package main

import (
	"fmt"
	"log"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)        // 遍历过的点
	loopVisited := make(map[string]bool) //移出图的点
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				// 开始深度遍历
				seen[item] = true
				visitAll(m[item])
				loopVisited[item] = true
				order = append(order, item)
			} else {
				if !loopVisited[item] {
					// 开始遍历但是已经移出图的点
					log.Fatal("loop")
				}
			}
		}
	}

	for k := range m {
		visitAll([]string{k})
	}

	return order
}
