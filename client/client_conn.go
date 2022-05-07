package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func sender(conn *websocket.Conn) error {
	for {
		if conn == nil {
			return errors.New("No web socket connection.")
		}
		fmt.Println("Sending message")
		message := []byte("MESSAGE SENT FROM MY CLIENT")
		_ = conn.WriteMessage(websocket.TextMessage, message)
		time.Sleep(2 * time.Second)
	}
}
