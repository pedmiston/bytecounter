package bytecounter

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
)

// A Total contains the summary statistics for a repo.
type Total struct {
	Name     string
	Forks    int
	Watching int
	Stars    int
	// Contributors int
	// Commits      int
	// Branches     int
	// Releases     int
	// Issues       int
	// PullRequests int
}

// Totals is a slice of Total objects.
type Totals []Total

// GetTotals returns the summary statistics for each repo in repos.
func GetTotals(repos []string) (Totals, error) {
	client := github.NewClient(nil)
	ctx := context.Background()

	totals := make(Totals, len(repos))
	for i, fullName := range repos {
		parts := strings.Split(fullName, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unable to parse repo: %v", fullName)
		}
		owner, name := parts[0], parts[1]
		repo, _, err := client.Repositories.Get(ctx, owner, name)
		if err != nil {
			return nil, fmt.Errorf("couldn't get repo %v: %v", fullName, err)
		}
		totals[i] = Total{
			Name:     repo.GetFullName(),
			Forks:    repo.GetForksCount(),
			Watching: repo.GetWatchersCount(),
			Stars:    repo.GetStargazersCount(),
		}

	}
	return totals, nil
}

// ToCSV converts a slice of Total objects to a slice of records
// ready to be written to CSV.
func (t Totals) ToCSV() [][]string {
	records := make([][]string, 1+len(t))
	records[0] = csvHeader
	for i, total := range t {
		records[i+1] = total.ToRecord()
	}
	return records
}

var csvHeader = []string{"Name", "Forks", "Watching", "Stars"}

// ToRecord converts the Total object to a slice of strings
func (t Total) ToRecord() []string {
	return []string{t.Name, strconv.Itoa(t.Forks), strconv.Itoa(t.Watching), strconv.Itoa(t.Stars)}
}
