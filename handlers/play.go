package handlers

import (
	"encoding/json"
	"fmt"
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
			log.Printf("error reading /play input: %s", err)
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}

		err = json.Unmarshal(body, &match)
		if err != nil {
			log.Printf("error unmarshalling play input: %s", err)
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}

		if match.Player < models.Rock || match.Player > models.Spock {
			log.Printf("player choice out of range: %s", match.Player)
			WriteError(w, http.StatusBadRequest, "player choice must be in range 1-5")
			return
		}

		computerChoice, err := GetRandomChoice()
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}
		match.Computer = computerChoice.ID

		match.DetermineWinner()

		m, err := json.Marshal(match)
		if err != nil {
			log.Printf("error marshalling play out: %s", err)
			WriteError(w, http.StatusInternalServerError, "internal error")
		}
		io.WriteString(w, string(m))
	default:
		e := fmt.Sprintf("given method not allowed: %s", r.Method)
		WriteError(w, http.StatusMethodNotAllowed, e)
	}
}
