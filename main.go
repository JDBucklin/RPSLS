package main

import (
	"fmt"
	"net/http"

	"github.com/jdbucklin/RPSLS/handlers"
)

func main() {
	http.HandleFunc("/choices", handlers.HandleChoices)
	http.HandleFunc("/choice", handlers.HandleChoice)
	http.HandleFunc("/play", handlers.HandlePlay)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
	}
}
