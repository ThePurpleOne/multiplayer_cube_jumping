package main

import rl "github.com/gen2brain/raylib-go/raylib"

type world struct {
	size rl.Vector2
	player_list []player
}


func create_world(w, h int32) world {
	// CREATE EMPTY WORLD WITH SIZE WxH
	return world{rl.NewVector2(float32(w), float32(h)), []player{}}
}

func (w world) dummy() {
	w.size.X = w.size.X
}


