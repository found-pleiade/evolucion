package main

import (
	"html/template"
	"io"
	"math/rand"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("./templates/*.*html")),
}

// Connection handles the connection of a player to the game.
// It checks if the player has already a session and if
// he is already in the game. If not, it saves the player to the game.
func connection(c echo.Context) error {
	session, _ := session.Get("session", c)
	if session != nil {
		if game.isPlayerPresent(session.Values["id"].(int)) {
			return c.Redirect(http.StatusFound, "/play")
		}
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// Check if name is empty
	name := c.FormValue("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Name is empty")
	}

	// Save player to game
	id := rand.Intn(1000)
	game.Players = append(game.Players, Player{ID: id, Name: name, Deck: []Card{carapace, carnivore}})

	// Save ID to session
	session.Values["id"] = id
	session.Values["name"] = name
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/play")
}

// Checks if player has already a session
// If yes, redirect to play page
// If not, redirect to login page
func index(c echo.Context) error {
	session, _ := session.Get("session", c)
	if session != nil {
		if game.isPlayerPresent(session.Values["id"].(int)) {
			return c.Redirect(http.StatusFound, "/play")
		}
	}

	return c.Render(http.StatusOK, "login", nil)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Renderer = t
	e.Static("/public", "public")     // Serve static files, like images
	e.GET("/play", Play)              // Play is the main page of the game
	e.GET("/", index)                 // Index is the login page
	e.POST("/connection", connection) // Connection handles the connection of a player to the game
	e.GET("/ws", ws)                  // ws is the websocket connection
	e.Logger.Fatal(e.Start(":1323"))
}
