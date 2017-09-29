package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/pedmiston/bytecounter"
)

var tmpl *template.Template

func main() {
	// Start a webserver to make the bytecounter package interactive
	var err error
	tmpl, err = template.ParseFiles("templates/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func Home(rw http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error

	switch r.Method {
	case "GET":
		// pass
	case "POST":
		query := r.FormValue("query")
		data, err = bytecounter.GetTotals([]string{query})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("POST %v: %v\n", query, data)
	}
	tmpl.Execute(rw, data)
}
