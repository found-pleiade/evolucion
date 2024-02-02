package main

import (
	"fmt"
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
	Phase struct {
		Description  string
		Name         string
		IsSequential bool
	}
	Player struct {
		ID      int
		Name    string
		Hand    []Card
		Species []Species
		IsReady bool
	}
	Species struct {
		BodySize   int
		Food       int
		Population int
		Name       string
		Traits     []Card
	}
)

func Play(c echo.Context) error {
	session, err := session.Get("session", c)
	if err != nil || session.Values["name"] == nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "game", session.Values["name"])
}

func (g *Game) InitializeGame() *Game {
	g.Deck = InitializeDeck()
	g.DiscardPile = []Card{}
	g.Food = Food{CurrentValue: 0, FutureValue: 0}
	g.Players = []Player{{ID: 55, Name: "Alexis", Hand: []Card{carapaceTemplate.Generate()[0], charognardTemplate.Generate()[0], longCouTemplate.Generate()[0]}}, {ID: 1050, Name: "Baptiste", Hand: []Card{chasseEnMeuteTemplate.Generate()[0]}}}
	return g
}

func (g *Game) RemovePlayer(id int) {
	for i, player := range g.Players {
		if player.ID == id {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
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
