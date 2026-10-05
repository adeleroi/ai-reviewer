package main

import (
	"ai-reviewer/internal/git"
	"ai-reviewer/internal/reviewer"
	"fmt"
)

func main() {
	diff, err := git.GetStateDiff("experiments/code-1.go")
	if err != nil {
		panic(err)
	}
	instruction := `You're an automated code reviewer. Analize the provided git diff
	highlight potential issues like correctness, security, reliability, performance, or meaningful maintainability
	in changes only
	`
	fmt.Println(diff)

	ctxFile, err := reviewer.ReadFile("./experiments/external-context-1.go")
	if err != nil {
		panic(err)
	}
	ctxStr := string(ctxFile)

	reviews, err := reviewer.GetReviewFor(diff, instruction, ctxStr)
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
