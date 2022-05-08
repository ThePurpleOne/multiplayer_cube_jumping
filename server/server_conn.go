package main

import (
	"fmt"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)


type wrapper_handler struct{
	id int32
	world_size rl.Vector2
	player_pos rl.Vector2
	color int32
	move string
}

// Handler fucntion called everytime there is a request on /ws
func (wh* wrapper_handler) handler_socket(w http.ResponseWriter, r *http.Request) { 
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// SEND INITIAL DATA TO CLIENT
	// 1 - Its ID
	// 3 - The World size
	// 3 - Its inital position
	msg := []byte(fmt.Sprintf("%d|%f,%f|%d|%f,%f", wh.id,
								wh.world_size.X,
								wh.world_size.Y,
								wh.color,
								wh.player_pos.X,
								wh.player_pos.Y))

	_ = conn.WriteMessage(websocket.TextMessage, msg)//msg)

	move_chan := make(chan string, 2)
	go reader(conn, move_chan)

	// After the first message, the wrapper can be modified to get the move back to the main
	wh.move = <-move_chan
}


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn* websocket.Conn, move_chan chan string) {
	for{
		_, payload, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(string(payload))
		move_chan <- string(payload)
		payload = []byte("Initial Connection")
		_ = conn.WriteMessage(websocket.TextMessage, payload)
	}
}

