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
		{
			"targets": []string{"http://example.com"},
			"labels":  map[string]string{},
		},
		{
			"targets": []string{"http://localhost:8000"},
			"labels":  map[string]string{},
		},
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
				"module":      "http_2xx",
				"application": "App 1",
				"team":        "Alpha",
				"environment": "prod",
				"datacenter":  "A",
			},
		},
		{
			"targets": []string{"http://localhost:8000"},
			"labels": map[string]string{
				"module":      "http_2xx",
				"application": "App 2",
				"team":        "Beta",
				"environment": "nonprod",
				"datacenter":  "B",
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
