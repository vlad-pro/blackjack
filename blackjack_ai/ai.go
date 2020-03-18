package blackjack_ai

import (
	"fmt"

	"github.com/gophercises/deck"
)

type AI interface {
	Play(hand [][]deck.Card, dealer deck.Card) Move
	Bet() int
	Results(hand []deck.Card, dealer []deck.Card)
}

type HumanAI struct{}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand [][]deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("Invaling option:", input)
		}
	}
}

func (ai *HumanAI) Results(hand []deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}

// Filler to comlete later
