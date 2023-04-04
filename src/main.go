package main

import (
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
			fmt.Println(name, value)
		}
	}
}
