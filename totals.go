package bytecounter

// A Total contains the summary statistics for a repo.
type Total struct {
	Name string
	// Commits      int
	// Branches     int
	// Releases     int
	// Contributors int
	// Watching     int
	// Stars        int
	// Forks        int
	// Issues       int
	// PullRequests int
}

// Totals returns the summary statistics for each repo in repos.
func Totals(repos []string) []Total {
	return nil
}
