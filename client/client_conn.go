package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

func int_to_color(i uint32) rl.Color {

	r := uint8((i >> 24) & 0xFF)
	g := uint8((i >> 16) & 0xFF)
	b := uint8((i >> 8)  & 0xFF)
	a := uint8((i) & 0xFF)

	return rl.NewColor(r, g, b, a)
	//return rl.NewColor(uint8(i & 0xFF), (uint8(i >> 8)), (uint8(i >> 16)), (uint8(i >> 24)))
}

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

func receive_init(conn *websocket.Conn) (int32, rl.Vector2, rl.Color, rl.Vector2) {
		_, payload, err := conn.ReadMessage()
		if err != nil{
			fmt.Println("Error: ", err)
		}
		fmt.Println("Received message: ", string(payload))
		splitted := strings.Split(string(payload), "|")

		// ! PARSE ID
		id, _ := strconv.ParseInt(splitted[0], 10, 32)
		log.Default().Println("ID: ", id)

		// ! PARSE SIZE
		size := strings.Split(splitted[1], ",")
		width, _ := strconv.ParseFloat(size[0], 64)
		height, _ := strconv.ParseFloat(size[1], 64)
		size_vec := rl.NewVector2(float32(width), float32(height))

		// ! PARSE COLOR
		c, _ := strconv.ParseInt(splitted[2], 10, 32)
				
		color := int_to_color(uint32(c))

		// ! PARSE POSITION
		pos := strings.Split(splitted[3], ",")
		pos_x, _ := strconv.ParseFloat(pos[0], 32)
		pos_y, _ := strconv.ParseFloat(pos[1], 32)
		pos_vec := rl.NewVector2(float32(pos_x), float32(pos_y))
		
		return 	int32(id), size_vec, color, pos_vec
}
