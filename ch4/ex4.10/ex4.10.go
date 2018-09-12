package main

import (
	"fmt"
	"gopl/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	lMonthSlice := make([]*github.Issue, 0)
	lYearSlice := make([]*github.Issue, 0)
	bYearSlice := make([]*github.Issue, 0)

	now := time.Now()
	// 一个月前
	lMonth := now.AddDate(0, -1, 0)
	// 一年前
	lYear := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(lMonth) {
			lMonthSlice = append(lMonthSlice, item)
		} else if item.CreatedAt.Before(lMonth) && item.CreatedAt.After(lYear) {
			lYearSlice = append(lYearSlice, item)
		} else {
			bYearSlice = append(bYearSlice, item)
		}
	}

	fmt.Printf("issue last one month\n")

	for _, issue := range lMonthSlice {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

	fmt.Printf("issue last one year\n")

	for _, issue := range lYearSlice {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}


	fmt.Printf("issue before one year\n")
	for _, issue := range bYearSlice {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title)
	}

}
