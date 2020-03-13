package models

// Match represents a single match of RPSLS
type Match struct {
	Results  string     `json:"results"`
	Player   ChoiceType `json:"player"`
	Computer ChoiceType `json:"computer"`
}

// Win, Lose, Tie constants
const (
	Win  string = "win"
	Lose string = "lose"
	Tie  string = "tie"
)

// DetermineWinner decides the winner of a game of RPSLS
func (m *Match) DetermineWinner() {
	switch m.Player {
	case Rock:
		switch m.Computer {
		case Scissors, Lizard:
			m.Results = Win
		case Paper, Spock:
			m.Results = Lose
		default:
			m.Results = Tie
		}
	case Paper:
		switch m.Computer {
		case Rock, Spock:
			m.Results = Win
		case Scissors, Lizard:
			m.Results = Lose
		default:
			m.Results = Tie
		}
	case Scissors:
		switch m.Computer {
		case Paper, Lizard:
			m.Results = Win
		case Rock, Spock:
			m.Results = Lose
		default:
			m.Results = Tie
		}
	case Lizard:
		switch m.Computer {
		case Paper, Spock:
			m.Results = Win
		case Rock, Scissors:
			m.Results = Lose
		default:
			m.Results = Tie
		}
	case Spock:
		switch m.Computer {
		case Rock, Scissors:
			m.Results = Win
		case Lizard, Paper:
			m.Results = Lose
		default:
			m.Results = Tie
		}
	}
}
