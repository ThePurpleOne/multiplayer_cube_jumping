package main

import (
	"fmt"
	"net/http"
)


const WIDTH = 800
const HEIGHT = 800

func main(){

	//p1 := create_player(100, 600, rl.Purple, 50);
	//w1 := create_world(WIDTH, HEIGHT);

	//p1.jump()
	//w1.dummy()

	fmt.Println("Starting server on port 8080")

	setup_routes()
	http.ListenAndServe(":8080", nil)

}
