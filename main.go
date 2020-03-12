package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/choices", HandleChoices)
	http.HandleFunc("/choice", HandleChoice)

	err := http.ListenAndServeTLS(port, "certificate.pem", "key.pem", nil)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
