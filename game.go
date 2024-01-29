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
	}
	Player struct {
		ID      int
		Name    string
		Deck    []Card
		Species []Species
	}
	Species struct {
		BodySize   int
		Food       int
		Population int
		Name       string
		Traits     []Card
	}
)

var game = Game{Food: Food{CurrentValue: 10, FutureValue: 19}, Players: []Player{{ID: 55, Name: "Alexis", Deck: []Card{carapaceTemplate.Gerenate()[0], charognardTemplate.Gerenate()[0], longCouTemplate.Gerenate()[0]}}, {ID: 1050, Name: "Baptiste", Deck: []Card{chasseEnMeuteTemplate.Gerenate()[0]}}}, Deck: []Card{carapaceTemplate.Gerenate()[0], longCouTemplate.Gerenate()[0]}, DiscardPile: []Card{cooperationTemplate.Gerenate()[0], cornesTemplate.Gerenate()[0]}}

func Play(c echo.Context) error {
	session, err := session.Get("session", c)
	if err != nil || session.Values["name"] == nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, "/")
	}

	return c.Render(http.StatusOK, "game", session.Values["name"])
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
