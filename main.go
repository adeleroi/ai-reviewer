package main

import (
	aireviewer "ai-reviewer/pkg/reviewer"
	"fmt"
)

func main() {
	diff, err := aireviewer.GetStateDiff("example-4.go")
	if err != nil {
		panic(err)
	}
	instruction := `You're an automated code reviewer. Analize the provided git diff
	highlight potential issues like correctness, security, reliability, performance, or meaningful maintainability
	in changes only
	`
	fmt.Println(diff)

	reviews, err := aireviewer.GetReviewFor([]byte(diff), instruction)
	if err != nil {
		panic(err)
	}
	for _, finding := range reviews.Findings {
		fmt.Println(finding.Title)
		fmt.Println(finding.File)
		fmt.Println(finding.Line)
		fmt.Println(finding.Explanation)
		fmt.Println(finding.Suggestion)
	}
}
