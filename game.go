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
		Name string
	}
	Player struct {
		ID   int
		Name string
		Deck []Card
	}

	Game struct {
		Food        Food
		Players     []Player
		Deck        []Card
		DiscardPile []Card
	}
)

var game = Game{Food: Food{CurrentValue: 10, FutureValue: 19}, Players: []Player{{ID: 55, Name: "Alexis", Deck: []Card{carapace, carapace, carnivore}}, {ID: 1050, Name: "Baptiste", Deck: []Card{carnivore}}}, Deck: []Card{carapace, carnivore}, DiscardPile: []Card{carapace}}

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
