package blackjack_ai

import "github.com/gophercises/deck"

type Game struct {
	// unexported fields
	deck   []deck.Card
	state  State
	player []deck.Card
	dealer []deck.Card
}

type Move func(Game) Game

func Hit(gs Game) Game {
	return gs
}

func Stand(gs Game) Game {
	return gs
}
