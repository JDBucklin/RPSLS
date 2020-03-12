package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Choice struct {
	ID   ChoiceType `json:"id"`
	Name string     `json:"name"`
}

func GetChoices() []Choice {
	return []Choice{
		Choice{Rock, Rock.String()},
		Choice{Paper, Paper.String()},
		Choice{Scissors, Scissors.String()},
		Choice{Lizard, Lizard.String()},
		Choice{Spock, Spock.String()},
	}
}

type ChoiceType int

const (
	Rock ChoiceType = iota + 1
	Paper
	Scissors
	Lizard
	Spock
)

func (c ChoiceType) String() string {
	switch c {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	case Lizard:
		return "Lizard"
	case Spock:
		return "Spock"
	}
	return ""
}

func HandleChoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	choices := GetChoices()
	c, err := json.Marshal(choices)
	if err != nil {
		//TODO
	}
	io.WriteString(w, string(c))
}

func HandleChoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	type Result struct {
		RandomNumber int `json:"random_number"`
	}

	resp, err := http.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		//TODO
	}
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//TODO
		}

		result := Result{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			//TODO
		}

		choiceType := ChoiceType(result.RandomNumber%5 + 1)
		choice := Choice{
			ID:   choiceType,
			Name: choiceType.String(),
		}
		c, err := json.Marshal(choice)
		if err != nil {
			//TODO
		}
		io.WriteString(w, string(c))
	}
}
