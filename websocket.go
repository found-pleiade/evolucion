package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Game      Game
	SessionID int
}

var upgrader = websocket.Upgrader{}

// ws is a function that handles the websocket connection.
func ws(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}

	defer ws.Close()
	isOpen := true

	go func() {
		for isOpen {
			var buf bytes.Buffer
			data := Data{Game: game, SessionID: session.Values["id"].(int)}
			err := t.Render(&buf, "board", data, c)
			if err != nil {
				fmt.Println(err)
			}
			err = ws.WriteMessage(websocket.TextMessage, buf.Bytes())
			if err != nil {
				if !websocket.IsCloseError(err, websocket.CloseGoingAway) {
					fmt.Println("write error:", err)
				}
				break
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for isOpen {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("read error:", err)
			}
			isOpen = false
			break
		}
		fmt.Println("Received from ws:")
		fmt.Printf("%s\n", msg)
	}
	return nil
}
