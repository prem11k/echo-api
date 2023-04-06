package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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
	clientName := ""
	for key, values := range w.Header() {
		headers[key] = values

		if strings.Contains(key, "Ssl-Client-Subject-Dn") {
			re := regexp.MustCompile(`CN=([^,]+)`)
			match := re.FindStringSubmatch(values[0])
			fmt.Println(match)
			if len(match) >= 2 {
				clientName = match[1]
			}
		}
	}

	response := struct {
		Headers map[string][]string `json:"headers"`
		Message string              `json:"message"`
	}{
		Headers: headers,
		Message: fmt.Sprintf("Hello, %s!", clientName),
	}

	jsonResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
