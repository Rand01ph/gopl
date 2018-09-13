package main

import (
	"fmt"
	"gopl/ch4/github"
	"log"
	"os"
)

func main() {

	result, err := github.CreateIssues(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.HTMLURL)

}

