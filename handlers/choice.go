package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/jdbucklin/RPSLS/models"
)

// HandleChoices handles the /choices endpoint
func HandleChoices(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		choices := models.GetChoices()
		c, err := json.Marshal(choices)
		if err != nil {
			log.Printf("error marshalling choices: %s", err)
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}
		io.WriteString(w, string(c))
	} else {
		e := fmt.Sprintf("given method not allowed: %s", r.Method)
		WriteError(w, http.StatusMethodNotAllowed, e)
	}
}

// HandleChoice handles the /choice endpoint
func HandleChoice(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		choice, err := GetRandomChoice()
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}
		c, err := json.Marshal(choice)
		if err != nil {
			log.Printf("error marshalling choice: %s", err)
			WriteError(w, http.StatusInternalServerError, "internal error")
			return
		}
		io.WriteString(w, string(c))
	} else {
		e := fmt.Sprintf("given method not allowed: %s", r.Method)
		WriteError(w, http.StatusMethodNotAllowed, e)
	}
}

// GetRandomChoice gets a random number from "https://codechallenge.boohma.com/random"
// and returns a random hand choice
func GetRandomChoice() (models.Choice, error) {
	type Result struct {
		RandomNumber int `json:"random_number"`
	}

	resp, err := http.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		log.Printf("error retrieving random number from https://codechallenge.boohma.com/random: %s", err)
	}

	choice := models.Choice{}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error reading response body from https://codechallenge.boohma.com/random: %s", err)
			return models.Choice{}, err
		}

		result := Result{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Printf("error unmarshalling random number: %s", err)
			return models.Choice{}, err
		}

		choiceType := models.ChoiceType(result.RandomNumber%5 + 1)
		choice.ID = choiceType
		choice.Name = choiceType.String()
	} else {
		// if the random number website doesn't work, then generate one locally
		choiceType := models.ChoiceType(rand.Intn(6))
		choice.ID = choiceType
		choice.Name = choiceType.String()
	}
	return choice, err
}
