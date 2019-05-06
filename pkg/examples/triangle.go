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
		test := visualgo.Map(800, vgo.Width)
		fmt.Println(test)
	}

	draw := func() {
		//fmt.Println("Draw Stuff Here")
		vgo.Triangle(vgo.Width/2, 1, vgo.Width, -1, 0, -1)
	}
	vgo.Init(setup, draw)

}
