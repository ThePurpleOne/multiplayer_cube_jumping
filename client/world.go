package main

import rl "github.com/gen2brain/raylib-go/raylib"

type world struct {
	size rl.Vector2
}

func create_world(w, h int32) world {
	return world{rl.NewVector2(float32(w), float32(h))}
}