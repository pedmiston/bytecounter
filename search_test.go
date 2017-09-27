package bytecounter

import "testing"

func TestSearchExact(t *testing.T) {
	query := "pedmiston/bytecounter"
	results, err := Search(query, 1)
	if err != nil {
		t.Fatal(err)
	}
	if r := results[0]; r != query {
		t.Errorf("searched for %s instead returned %s", query, r)
	}
}

func TestSearchLimit(t *testing.T) {
	query := "node"
	limit := 10
	results, err := Search(query, limit)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != limit {
		t.Errorf("should %v results, instead got %v", limit, len(results))
	}
}

func TestSearchShrink(t *testing.T) {
	query := "pedmiston/bytecounter"
	results, err := Search(query, 100)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 {
		t.Errorf("more results than expected")
	}
}
