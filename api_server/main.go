package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
}

var people = []Person{
	{ID: 1, Name: "Alice", Email: "alice@example.com", Phone: "123-456-7890", Bio: "Engineer from NY."},
	{ID: 2, Name: "Bob", Email: "bob@example.com", Phone: "234-567-8901", Bio: "Designer from SF."},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com", Phone: "345-678-9012", Bio: "Manager from TX."},
	{ID: 4, Name: "Diana", Email: "diana@example.com", Phone: "456-789-0123", Bio: "Product manager from WA."},
	{ID: 5, Name: "Ethan", Email: "ethan@example.com", Phone: "567-890-1234", Bio: "Data scientist from MA."},
	{ID: 6, Name: "Fiona", Email: "fiona@example.com", Phone: "678-901-2345", Bio: "Marketing expert from IL."},
	{ID: 7, Name: "George", Email: "george@example.com", Phone: "789-012-3456", Bio: "Sales rep from FL."},
	{ID: 8, Name: "Hannah", Email: "hannah@example.com", Phone: "890-123-4567", Bio: "UX researcher from CO."},
	{ID: 9, Name: "Ian", Email: "ian@example.com", Phone: "901-234-5678", Bio: "DevOps engineer from AZ."},
	{ID: 10, Name: "Jasmine", Email: "jasmine@example.com", Phone: "012-345-6789", Bio: "HR specialist from GA."},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/data", handleList)
	r.Get("/data/{id}", handleDetail)

	log.Println("API server running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func handleList(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.URL.Query().Get("search"))
	var filtered []Person
	for _, p := range people {
		if search == "" || strings.Contains(strings.ToLower(p.Name), search) {
			filtered = append(filtered, p)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func handleDetail(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, p := range people {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.NotFound(w, r)
}
