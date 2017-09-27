package bytecounter

import "testing"

func TestSearchExactName(t *testing.T) {
	var results []string
	name := "pedmiston/bytecounter"
	results = Search(name, 1)
	if len(results) == 0 {
		t.Fatalf("search for '%s' returned no results", name)
	}
	if r := results[0]; r != name {
		t.Errorf("searched for %s instead returned %s", name, r)
	}
}
