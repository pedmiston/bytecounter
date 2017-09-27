package bytecounter

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// Search GitHub and return repo names as results.
func Search(query string, limit int) ([]string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	result, _, err := client.Search.Repositories(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	if len(result.Repositories) == 0 {
		return nil, fmt.Errorf("query '%s' returned no results", query)
	}

	nResults := min(len(result.Repositories), limit)
	names := make([]string, nResults)
	for i := 0; i < nResults; i++ {
		names[i] = result.Repositories[i].GetFullName()
	}
	return names, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
