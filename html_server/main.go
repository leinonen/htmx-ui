package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"coolbeans/pkg/api"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", indexHandler)
	r.Get("/search", searchHandler)
	r.Get("/person/{id}", detailHandler)

	log.Println("HTML server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
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
	id := chi.URLParam(r, "id")
	person, err := api.GetPersonByID(id)
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
