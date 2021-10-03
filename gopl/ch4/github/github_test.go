package github

import (
	"fmt"
	"log"
	"testing"
)

func TestSearchIssues(t *testing.T) {
	testQueryParams := []string{"repo:golang/go", "is:open", "json", "decode"}
	result, err := SearchIssues(testQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}