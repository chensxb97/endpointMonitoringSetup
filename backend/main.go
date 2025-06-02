package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome! I am a backend service for illustrating HTTP Service Discovery.")
}

func targetsHandler(w http.ResponseWriter, r *http.Request) {
	targets := []map[string]interface{}{
		{"targets": []string{"http://example.com"}},
		{"targets": []string{"http://localhost:8000"}},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}

func targetsHandlerWithCustomLabels(w http.ResponseWriter, r *http.Request) {
	targets := []map[string]interface{}{
		{
			"targets": []string{"http://example.com"},
			"labels": map[string]string{
				"__meta_module":      "http_2xx",
				"__meta_application": "App 1",
				"__meta_team":        "Alpha",
				"__meta_environment": "prod",
				"__meta_datacenter":  "A",
			},
		},
		{
			"targets": []string{"http://localhost:8000"},
			"labels": map[string]string{
				"__meta_module":      "http_2xx",
				"__meta_application": "App 2",
				"__meta_team":        "Beta",
				"__meta_environment": "nonprod",
				"__meta_datacenter":  "B",
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/targets", targetsHandler)
	http.HandleFunc("/targets_custom", targetsHandlerWithCustomLabels)

	protocol := "http"
	host := "localhost"
	port := ":8000"
	fmt.Printf("Starting server on %s://%s%s\n", protocol, host, port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
