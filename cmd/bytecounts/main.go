package main

import (
	"flag"
	"os"

	"github.com/pedmiston/bytecounter"
)

func main() {
	query := flag.String("q", "", "term to search for on github.com")
	limit := flag.Int("l", 10, "only get this many results")
	flag.Parse()
	repos := append(flag.Args(), bytecounter.Search(*query, *limit)...)
	os.Stdout.Write(bytecounter.Totals(repos))
}
