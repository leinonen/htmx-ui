package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Bio   string `json:"bio"`
}

var ErrNotFound = errors.New("resource not found")
var ErrDecode = errors.New("failed to decode data")

func SearchPerson(search string) []Person {
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

func GetPerson(r *http.Request) (*Person, error) {
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
