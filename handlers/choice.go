package handlers

import (
	"encoding/json"
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
			w.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(w, string(c))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// HandleChoice handles the /choice endpoint
func HandleChoice(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		choice, err := GetRandomChoice()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		c, err := json.Marshal(choice)
		if err != nil {
			log.Printf("error marshalling choice: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		io.WriteString(w, string(c))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		// if the random number website doesn't work generate one locally
		choiceType := models.ChoiceType(rand.Intn(6))
		choice.ID = choiceType
		choice.Name = choiceType.String()
	}
	return choice, err
}
