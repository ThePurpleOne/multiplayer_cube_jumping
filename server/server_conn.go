package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn* websocket.Conn){
	for{
		messageType, payload, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(string(payload))
		payload = append(payload, "ADDED FROM THE SERVER"...)
		_ = conn.WriteMessage(messageType, payload)
	}
}

func handle_index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handler_socket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader(conn)
}

func setup_routes(){
	http.HandleFunc("/", handle_index)
	http.HandleFunc("/ws", handler_socket)
}