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
		{"targets": []string{"http://example.com2"}},
		{"targets": []string{"https://google.com"}},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(targets); err != nil {
		http.Error(w, "Failed to encode targets", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/targets", targetsHandler)

	protocol := "http"
	host := "localhost"
	port := ":8000"
	fmt.Printf("Starting server on %s://%s%s\n", protocol, host, port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
