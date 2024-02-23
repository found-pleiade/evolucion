package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type (
	Food struct {
		CurrentValue int
		FutureValue  int
	}
	Card struct {
		Name             string
		ShortDescription string
		LongDescription  string
		Color            string
		IsPrior          bool
		Carnivore        bool
		FoodPoints       int
	}
	Game struct {
		Food        Food
		Players     []Player
		Deck        []Card
		DiscardPile []Card
		Phase       Phase
		PlayerTurn  int // Index of the player in the Players slice
	}
)

func play(c echo.Context) error {
	session, err := session.Get("session", c)
	if err != nil || session.Values["name"] == nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "game", session.Values["name"])
}

func (g *Game) initialize() {
	g.Deck = initializeDeck()
	g.Phase = gamePhases[PhaseWait]
	// For dev purposes
	g.mock()
}

func (g *Game) addPlayer(name string) int {
	var newPlayer Player
	newPlayer.initialize(rand.Intn(1000), name)

	g.Players = append(g.Players, newPlayer)

	return newPlayer.ID
}

func (g *Game) removePlayer(id int) {
	for i, player := range g.Players {
		if player.ID == id {
			g.DiscardPile = append(g.DiscardPile, g.Players[i].Hand...)
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			break
		}
	}
}

func (g *Game) isPlayerPresent(id int) bool {
	if len(g.Players) > 0 {
		for _, player := range g.Players {
			if player.ID == id {
				return true
			}
		}
	}

	return false
}

func (g *Game) mock() {
	g.Players = []Player{{ID: 55, Name: "Alexis", Hand: []Card{carapaceTemplate.generate()[0], charognardTemplate.generate()[0], longCouTemplate.generate()[0]}}, {ID: 1050, Name: "Baptiste", Hand: []Card{chasseEnMeuteTemplate.generate()[0]}}}
}
