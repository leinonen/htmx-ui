package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"coolbeans/pkg/api"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/person/", detailHandler)

	log.Println("HTML server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	items := api.SearchPerson("")
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/header.html",
		"templates/sidebar.html",
		"templates/index.html",
		"templates/results.html",
	))
	tmpl.ExecuteTemplate(w, "base", items)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	items := api.SearchPerson(search)

	tmpl := template.Must(template.ParseFiles("templates/results.html"))
	tmpl.Execute(w, items)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	person, err := api.GetPerson(r)
	if err != nil {
		if errors.Is(err, api.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/header.html",
		"templates/sidebar.html",
		"templates/detail.html",
	))
	tmpl.ExecuteTemplate(w, "base", person)
}
