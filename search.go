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

	names := make([]string, limit)
	for i, repository := range result.Repositories {
		names[i] = repository.GetFullName()
		if i == limit-1 {
			break
		}
	}
	return names, nil
}
