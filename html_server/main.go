package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var ErrNotFound = errors.New("resource not found")
var ErrDecode = errors.New("failed to decode data")

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/person/", detailHandler)

	log.Println("HTML server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	items := fetchData("")
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
	items := fetchData(search)

	tmpl := template.Must(template.ParseFiles("templates/results.html"))
	tmpl.Execute(w, items)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	person, err := fetchPerson(r)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
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

func fetchData(search string) []Person {
	apiURL := getApiUrl()
	resp, err := http.Get(apiURL + "/data?search=" + search)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var items []Person
	json.NewDecoder(resp.Body).Decode(&items)
	return items
}

func fetchPerson(r *http.Request) (*Person, error) {
	apiUrl := getApiUrl()
	idStr := strings.TrimPrefix(r.URL.Path, "/person/")
	resp, err := http.Get(apiUrl + "/data/" + idStr)
	if err != nil || resp.StatusCode != 200 {
		return nil, ErrNotFound
	}
	defer resp.Body.Close()

	var person Person
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		return nil, ErrDecode
	}
	return &person, nil
}

func getApiUrl() string {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8081" // fallback for local dev
	}
	return apiURL
}
