package main

import (
	"flag"
	"fmt"
	"gopl/ch4/github"
)

func main() {

	var opType = flag.String("t", "create", "选择操作类型")

	flag.Parse()
	fmt.Printf("opType is %s", *opType)

	switch *opType {
	case "create":
		github.CreateIssues(flag.Arg(0), flag.Arg(1), flag.Arg(2))
	case "read":
		issue, err := github.GetIssues(flag.Arg(0), flag.Arg(1), flag.Arg(2))
		if err != nil {
			fmt.Printf("read issues err is %s", err)
		}
		fmt.Printf("issue titile is %s body is %s\n", issue.Title, issue.Body)
	case "update":
		issue, err := github.UpdateIssues(flag.Arg(0), flag.Arg(1), flag.Arg(2))
		if err != nil {
			fmt.Printf("update issues err is %s", err)
		}
		fmt.Printf("issue titile is %s body is %s\n", issue.Title, issue.Body)
	case "close":
		status, err := github.CloseIssues(flag.Arg(0), flag.Arg(1), flag.Arg(2))
		if err != nil {
			fmt.Printf("close issues err is %s", err)
		}
		fmt.Printf("close issues status is %s", status)
	}
}
