package bytecounter

import "testing"

func TestSearchExactName(t *testing.T) {
	var results []string
	var err error
	name := "pedmiston/bytecounter"
	results, err = Search(name, 1)
	if err != nil {
		t.Fatal(err)
	}
	if r := results[0]; r != name {
		t.Errorf("searched for %s instead returned %s", name, r)
	}
}
