package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type (
	Card struct {
		Name string `json:"name" xml:"name" form:"name" query:"name"`
	}
	Player struct {
		ID   int    `json:"id" xml:"id" form:"id" query:"id"`
		Name string `json:"name" xml:"name" form:"name" query:"name"`
		Deck []Card `json:"deck" xml:"deck" form:"deck" query:"deck"`
	}
	Game struct {
		Players     []Player `json:"players" xml:"players" form:"players" query:"players"`
		Deck        []Card   `json:"deck" xml:"deck" form:"deck" query:"deck"`
		DiscardPile []Card
	}
)

var game = Game{Players: []Player{{ID: 55, Name: "Alexis", Deck: []Card{carapace, carapace, carnivore}}, {ID: 1050, Name: "Baptiste", Deck: []Card{carnivore}}}, Deck: []Card{carapace, carnivore}, DiscardPile: []Card{carapace}}

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
