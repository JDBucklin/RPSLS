package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jdbucklin/RPSLS/models"
)

// WriteError writes the given error status and detail to the response writer
func WriteError(w http.ResponseWriter, status int, detail string) {
	// write status header
	w.WriteHeader(status)

	// prepare json response
	e := models.ErrorStatus{Status: status, Detail: detail}
	es, err := json.Marshal(e)
	if err != nil {
		log.Printf("error marshalling error status: %s", err)
		// if there's an error just write out the detail as a string and return
		io.WriteString(w, detail)
		return
	}
	io.WriteString(w, string(es))
}
