package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TPlayerStream struct {
	PlayerId      string
	Message       chan string
	NewClients    chan chan string
	ClosedClients chan chan string
	TotalClients  map[chan string]bool
}
type TClientChan chan string

func (stream *TPlayerStream) listen() {
	for {
		select {
		// Add new available client
		case newClient := <-stream.NewClients:
			stream.TotalClients[newClient] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case closedClient := <-stream.ClosedClients:
			delete(stream.TotalClients, closedClient)
			close(closedClient)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *TPlayerStream) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(TClientChan)

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}
