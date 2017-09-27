package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"

	"github.com/pedmiston/bytecounter"
)

func main() {
	query := flag.String("q", "", "term to search for on github.com")
	limit := flag.Int("l", 10, "only get this many results")
	flag.Parse()
	repos := flag.Args()
	if *query != "" {
		found, err := bytecounter.Search(*query, *limit)
		if err != nil {
			log.Fatal(err)
		}
		repos = append(repos, found...)
	}
	totals, err := bytecounter.GetTotals(repos)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(os.Stdout)
	writer.WriteAll(totals.ToCSV())
}
