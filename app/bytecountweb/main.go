package main

import (
	"net/http"
)

func main() {
	// Start a webserver to make the bytecounter package interactive
	http.HandleFunc("/search", Search)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func Search(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(r.FormValue("query")))
}
