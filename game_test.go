package main

import (
	"testing"
)

// TestRemovePlayerNominal tests the nominal case where a player is removed from the game.
func TestRemovePlayerNominal(t *testing.T) {
	game := Game{Players: []Player{{ID: 1}, {ID: 2}, {ID: 3}}}
	want := Game{Players: []Player{{ID: 1}, {ID: 3}}}
	game.RemovePlayer(2)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

// TestRemovePlayerNonPresent tests the case where the id of the player to remove does not exist.
func TestRemovePlayerNonPresent(t *testing.T) {
	game := Game{Players: []Player{{ID: 1}, {ID: 2}, {ID: 3}}}
	want := Game{Players: []Player{{ID: 1}, {ID: 2}, {ID: 3}}}
	game.RemovePlayer(4)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

// TestRemovePlayerZeroPlayer tests the case where a player is removed from the game and there are no players left.
func TestRemovePlayerZeroPlayer(t *testing.T) {
	game := Game{Players: []Player{{ID: 1}}}
	want := Game{Players: []Player{}}
	game.RemovePlayer(1)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

// TestRemovePlayerNilPlayer tests the case where their is no player to remove.
func TestRemovePlayerNilPlayer(t *testing.T) {
	game := Game{Players: []Player{}}
	want := Game{Players: []Player{}}
	game.RemovePlayer(1)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

// TestRemovePlayerDiscardPile tests the case where a player is removed from the game and their hand is added to the discard pile.
func TestRemovePlayerDiscardPile(t *testing.T) {
	game := Game{Players: []Player{{ID: 1, Hand: []Card{{Name: "test1"}}}, {ID: 2, Hand: []Card{{Name: "test2"}}}, {ID: 3, Hand: []Card{{Name: "test3"}}}}, DiscardPile: []Card{{Name: "test4"}}}
	want := Game{Players: []Player{{ID: 1, Hand: []Card{{Name: "test1"}}}, {ID: 3, Hand: []Card{{Name: "test3"}}}}, DiscardPile: []Card{{Name: "test4"}, {Name: "test2"}}}
	game.RemovePlayer(2)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

// TestRemovePlayerDiscardPileWithoutHand tests the case where a player has no hand and is removed from the game.
func TestRemovePlayerDiscardPileWithoutHand(t *testing.T) {
	game := Game{Players: []Player{{ID: 1, Hand: []Card{{Name: "test1"}}}, {ID: 2}, {ID: 3, Hand: []Card{{Name: "test3"}}}}, DiscardPile: []Card{{Name: "test4"}}}
	want := Game{Players: []Player{{ID: 1, Hand: []Card{{Name: "test1"}}}, {ID: 3, Hand: []Card{{Name: "test3"}}}}, DiscardPile: []Card{{Name: "test4"}}}
	game.RemovePlayer(2)
	if !game.Equal(want) {
		t.Errorf("RemovePlayer(): got %v, want %v", game, want)
	}
}

func TestIsPlayerPresentTrue(t *testing.T) {
	game := Game{Players: []Player{{ID: 1}, {ID: 2}, {ID: 3}}}
	if !game.isPlayerPresent(2) {
		t.Errorf("isPlayerPresent(): Player with ID 2 should exist in %v", game)
	}
}

func TestIsPlayerPresentFalse(t *testing.T) {
	game := Game{Players: []Player{{ID: 1}, {ID: 2}, {ID: 3}}}
	if game.isPlayerPresent(4) {
		t.Errorf("isPlayerPresent(): Player with ID 4 should not exist in %v", game)
	}
}

// Equal checks if two Game instances are equal.
func (g Game) Equal(other Game) bool {
	if len(g.Players) != len(other.Players) && len(g.DiscardPile) != len(other.DiscardPile) {
		return false
	}
	for i := range g.Players {
		if g.Players[i].ID != other.Players[i].ID {
			return false
		}
	}
	for i := range g.DiscardPile {
		if g.DiscardPile[i].Name != other.DiscardPile[i].Name {
			return false
		}
	}
	return true
}
