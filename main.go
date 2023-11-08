package main

import (
	"fmt"

	"f3s.tech/f3s-jira/pkg/f3sjira"
	"github.com/andygrunwald/go-jira"
)

func main() {
	config := f3sjira.JiraConfig{
		Url:  "https://issues.apache.org/jira/",
		Auth: nil,
	}

	client := f3sjira.NewJiraClient(config)

	client.Connect()
	client.Search("project = \"Apache Blur\" order by key asc", f3sjira.NewDefaultSearchOptions(), printIssue)
}

func printIssue(issue jira.Issue) error {
	fmt.Println(issue.Key, "; ", issue.Fields.TimeOriginalEstimate/3600, "h ; ", issue.Fields.Summary)
	return nil
}
