package main

import (
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

	go sender(conn)
	//time.Sleep(time.Second * 10)

	//// ! --------------------------------------------------
	//// ! ------------------ WORLD CREATIO------------------
	//// ! --------------------------------------------------
	p1 := create_player(100, 600, rl.Purple, 50);
	w1 := create_world(WIDTH, HEIGHT);

	rl.InitWindow(WIDTH, HEIGHT, "raylib [core] example - basic window")
	rl.SetTargetFPS(100)

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(WIDTH-100, 10)

		// ! --------------------------------------------------
		// ! -------------------- UPDATE -------------------
		// ! --------------------------------------------------
		if rl.IsKeyDown(rl.KeyLeft) && p1.check_left(w1) {
			p1.move_left(STEP);
		}
		if rl.IsKeyDown(rl.KeyRight) && p1.check_right(w1) {
			p1.move_right(STEP);
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			p1.is_jumping = true;
		}
	//	//else{
	//	//	// 
	//	//	if rl.IsKeyDown(rl.KeyUp) && p1.check_top(w1) {
	//	//		p1.move_up(STEP);
	//	//	}
	//	//	if rl.IsKeyDown(rl.KeyDown) && p1.check_bot(w1) {
	//	//		p1.move_down(STEP);
	//	//	}
	//	//}
		p1.jump()

		// ! --------------------------------------------------
		// ! -------------------- DRAWING -------------------
		// ! --------------------------------------------------
		p1.draw()
		rl.DrawText("CUBES", 10, 10, 30, rl.Blue)


		rl.EndDrawing()
	}
	rl.CloseWindow()
}