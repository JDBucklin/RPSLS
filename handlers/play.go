package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jdbucklin/RPSLS/models"
)

// HandlePlay handles the /play endpoint
func HandlePlay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case http.MethodOptions:
	case http.MethodPost:
		match := models.Match{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = json.Unmarshal(body, &match)
		if err != nil {
			log.Printf("error unmarshalling play input: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		computerChoice, err := GetRandomChoice()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		match.Computer = computerChoice.ID

		match.DetermineWinner()

		m, err := json.Marshal(match)
		if err != nil {
			log.Printf("error marshalling play out: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(w, string(m))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
