package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPerson(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data/1", nil)
	w := httptest.NewRecorder()

	handleDetail(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var p Person
	json.NewDecoder(resp.Body).Decode(&p)
	if p.Name != "Alice" {
		t.Errorf("expected Alice, got %s", p.Name)
	}
}

func TestSearchPerson(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data?search=bob", nil)
	w := httptest.NewRecorder()

	handleList(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var people []Person
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(people) != 1 {
		t.Fatalf("expected 1 result, got %d", len(people))
	}

	if people[0].Name != "Bob" {
		t.Errorf("expected Bob, got %s", people[0].Name)
	}
}
