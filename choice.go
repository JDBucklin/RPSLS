package main

import (
	"encoding/json"
	"io"
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

func Choices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	choices := GetChoices()
	c, _ := json.Marshal(choices)
	io.WriteString(w, string(c))
}
