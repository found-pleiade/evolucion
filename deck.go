package main

import (
	"math/rand"
)

func initializeDeck() []Card {
	var deck []Card
	for _, ct := range cardTemplates {
		deck = append(deck, ct.generate()...)
	}
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}
