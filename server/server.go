package main

import (
	"fmt"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 800


// Transform Color into int (RRRRRRRR GGGGGGGG BBBBBBBB AAAAAAAA)
func color_to_int(c rl.Color) int32 {
	return int32((int32(c.R) << 24) | (int32(c.G) << 16) | (int32(c.B) << 8) | int32(c.A))
	//return uint32(c.R)<<24 | uint32(c.G)<<16 | uint32(c.B)<<8 | uint32(c.A)
	//return (uint32(c.R) << 24) + (uint32(c.G) << 16) + (uint32(c.B) << 8) + (uint32(c.A))
}


func main(){

	// CREATE WORLD
	w1 := create_world(WIDTH, HEIGHT);

	// WRAPPER FOR HANDLER TO PASS INITIAL DATA  
	wh1 := wrapper_handler{	id: 0,
							world_size: w1.size,
							player_pos: rl.NewVector2(float32(WIDTH/2), float32(600)),
							color     : color_to_int(rl.Purple)}

	// ! ADD A PLAYER
	w1.player_list = append(w1.player_list, create_player(&wh1.id, 100, 600, rl.Purple, 50));

	

	//w1.player_list[0].

	// CREATE LISTENNER SERVER
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/ws", wh1.handler_socket)
	http.ListenAndServe(":8080", nil)

	fmt.Println("PLAYER MOVE : " + wh1.move) 
}
