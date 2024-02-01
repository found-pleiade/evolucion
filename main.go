package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/gin-gonic/gin"
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
	session, err := session.Get("session", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if session.Values["id"] != nil {
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
	game.Players = append(game.Players, Player{ID: id, Name: name, Hand: []Card{carapaceTemplate.Gerenate()[0], cornesTemplate.Gerenate()[0]}})

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
	session, err := session.Get("session", c)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	if session.Values["id"] != nil {
		if game.isPlayerPresent(session.Values["id"].(int)) {
			return c.Redirect(http.StatusFound, "/play")
		}
	}

	return c.Render(http.StatusOK, "login", nil)
}

func mainbis() {
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

// Initialize event and Start procnteessing requests
func NewPlayerServer(playerId string) (event *TPlayerStream) {
	event = &TPlayerStream{
		PlayerId:      playerId,
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}
	go event.listen()
	return
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "text/event-stream")
		context.Writer.Header().Set("Cache-Control", "no-cache")
		context.Writer.Header().Set("Connection", "keep-alive")
		context.Writer.Header().Set("Transfer-Encoding", "chunked")
		context.Next()
	}
}

func timeMessageLoop(stream *TPlayerStream) {
	for {
		time.Sleep(time.Second * 10)
		now := time.Now().Format("2006-01-02 15:04:05")
		currentTime := fmt.Sprintf("The Current Time Is %v", now)

		// Send current time to clients message channel
		stream.Message <- currentTime
	}
}

func main() {
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.GET("/connection", func(context *gin.Context) {
		log.Print("test connexion")
		context.Next()
	})

	stream := NewPlayerServer("test")
	go timeMessageLoop(stream)

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{"admin": "admin123"}))
	authorized.GET("/stream", HeadersMiddleware(), stream.serveHTTP(),
		func(context *gin.Context) {
			getValue, ok := context.Get("clientChan")
			if !ok {
				return
			}
			clientChan, ok := getValue.(TClientChan)
			if !ok {
				return
			}
			context.Stream(func(w io.Writer) bool {
				// Stream message to client from message channel
				if msg, ok := <-clientChan; ok {
					context.SSEvent("message", msg)
					return true
				}
				return false
			})
		},
	)

	router.Run(":1323")
}
