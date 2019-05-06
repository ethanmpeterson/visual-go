package main

import "fmt"
import "../visualgo"

func main() {
	fmt.Println("TODO")

	vgo := visualgo.WindowConfig{
		Width:  800,
		Height: 800,
		Title:  "Triangle",
	}

	setup := func() {
		fmt.Println("Visual Go Triangle Example")
	}

	draw := func() {
		//fmt.Println("Draw Stuff Here")
		//vgo.Triangle(vgo.Width/2, 0, vgo.Width, vgo.Height, 0, vgo.Height)
		vgo.Rect(100, 100, 400, 400)
	}
	vgo.Init(setup, draw)

}
