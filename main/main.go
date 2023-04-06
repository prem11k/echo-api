package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting on port 3000")
	http.HandleFunc("/", handleQuery)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reponse: ", w)
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(http.StatusOK)

	headers := make(map[string][]string)
	for key, values := range w.Header() {
		headers[key] = values
	}

	response := struct {
		Headers map[string][]string `json:"headers"`
		Message string              `json:"message"`
	}{
		Headers: headers,
		Message: "Hello, World!",
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
