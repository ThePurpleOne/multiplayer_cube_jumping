package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type player struct{
	pos rl.Vector2
	vel rl.Vector2
	color rl.Color
	size rl.Vector2
	is_jumping bool
	jump_count int32
}

const STEP = 10
func create_player(x_, y_ float32, c_ rl.Color, s_ float32) player {
	return player{	pos : rl.NewVector2(x_, y_),
					vel : rl.NewVector2(10, 10),
					color : c_,
					size : rl.NewVector2(s_, s_),
					is_jumping : false,
					jump_count : 10}
}

func (p player) draw(){
	rl.DrawRectangleV(p.pos, p.size, p.color)
}

// ! --------------------------------------------------
// ! -------------------- MOVEMENTS -------------------
// ! --------------------------------------------------

func (p* player) move_up(step int32){
	p.pos.Y -= p.vel.Y
}

func (p* player) move_down(step int32){
	p.pos.Y += p.vel.Y
}

func (p* player) move_left(step int32){
	p.pos.X -= p.vel.X
}

func (p* player) move_right(step int32){
	p.pos.X += p.vel.X
}

// CHECK IF THE NEXT ITERATION HITS A WALL 
func (p* player) check_top(w world) bool{
	return (p.pos.Y - p.vel.Y) >= 0
}
func (p* player) check_bot(w world) bool{
	return ((p.pos.Y + p.vel.Y) + p.size.Y) <= w.size.Y
}
func (p* player) check_left(w world) bool{
	return (p.pos.X - p.vel.X) >= 0
}
func (p* player) check_right(w world) bool{
	// Check if a player is going to hit the right wall of the world
	return ((p.pos.X + p.vel.X) + p.size.X) <= w.size.X
}

func (p* player) jump(){
	// source : https://stackoverflow.com/questions/51460626/how-to-make-an-object-jump-in-pygame#:~:text=you%20can%20only%20begin%20a,KEYDOWN)%20not%20if%20pressed.
	if p.is_jumping{
		if p.jump_count >= -10{
			neg := float64(1)
			if p.jump_count < 0{
				neg = -1
			}
			p.pos.Y -= float32(math.Pow(float64(p.jump_count), 2.0) * neg) * 0.7
			p.jump_count--
		}else{
			p.is_jumping = false
			p.jump_count = 10
		}
	}
}