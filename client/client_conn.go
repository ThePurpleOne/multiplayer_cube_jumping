package main

import (
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
)

func sender(conn *websocket.Conn, msg string) error {
	//for {
	if conn == nil {
		return errors.New("no web socket connection")
	}
	fmt.Println("Sending message")
	message := []byte(msg)
	_ = conn.WriteMessage(websocket.TextMessage, message)
	return nil
	//time.Sleep(2 * time.Second)
	//}
}
