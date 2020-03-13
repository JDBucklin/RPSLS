package models

// Choice holds the id and string for a given hand in RPSLS
type Choice struct {
	ID   ChoiceType `json:"id"`
	Name string     `json:"name"`
}

// GetChoices returns all choices
func GetChoices() []Choice {
	return []Choice{
		Choice{Rock, Rock.String()},
		Choice{Paper, Paper.String()},
		Choice{Scissors, Scissors.String()},
		Choice{Lizard, Lizard.String()},
		Choice{Spock, Spock.String()},
	}
}

// ChoiceType is the integer value of the chosen hand in RPSLS
type ChoiceType int

// ChoiceType constants
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
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	case Lizard:
		return "lizard"
	case Spock:
		return "spock"
	}
	return ""
}
