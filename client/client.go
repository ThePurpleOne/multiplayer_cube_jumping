package main

import (
	"fmt"
	"log"
	"net/url"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gorilla/websocket"
)

const WIDTH = 800
const HEIGHT = 800

func main() {

	u, _ := url.Parse("ws://localhost:8080/ws")
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil);
	if err != nil {
		//log.Printf("Handshake failed : %d", resp.StatusCode)
		log.Printf("Handshake failed : ")
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// ! RECEIVE INITIALIZATION DATA
	id , world_size, cube_color, player_pos := receive_init(conn)

	fmt.Println("ID: ", id)
	fmt.Println("World size: ", world_size)
	fmt.Println("Cube color: ", cube_color)
	fmt.Println("Player pos: ", player_pos)

	// ! TEST SENDING DATA
	sender(conn, "Hello0.")


	// ! --------------------------------------------------
	// ! ------------------ WORLD CREATIO------------------
	// ! --------------------------------------------------
	p1 := create_player(player_pos.X, player_pos.Y, cube_color, 50);
	//p1 := create_player(player_pos.X, player_pos.Y, rl.Purple, 50);
	w1 := create_world(int32(world_size.X), int32(world_size.Y));

	rl.InitWindow(WIDTH, HEIGHT, "raylib [core] example - basic window")
	rl.SetTargetFPS(100)

	for !rl.WindowShouldClose(){

		// ! --------------------------------------------------
		// ! -------------------- UPDATE ----------------------
		// ! --------------------------------------------------
		if rl.IsKeyDown(rl.KeyLeft) && p1.check_left(w1) {
			sender(conn, "L");
			p1.move_left(STEP);
		}
		if rl.IsKeyDown(rl.KeyRight) && p1.check_right(w1) {
			sender(conn, "R");
			p1.move_right(STEP);
		}

		if rl.IsKeyPressed(rl.KeySpace) && !p1.is_jumping{
			sender(conn, "J");
			p1.is_jumping = true;
		}else{
			if rl.IsKeyDown(rl.KeyUp) && p1.check_top(w1) {
				p1.move_up(STEP);
			}
			if rl.IsKeyDown(rl.KeyDown) && p1.check_bot(w1) {
				p1.move_down(STEP);
			}
		}
		p1.jump()
		
		// Debug p1 pos
		//fmt.Println("Player pos: ", p1.pos)

		// ! --------------------------------------------------
		// ! -------------------- DRAWING ---------------------
		// ! --------------------------------------------------
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(WIDTH-100, 10)
		p1.draw()
		rl.DrawText("CUBES", 10, 10, 30, rl.Blue)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}