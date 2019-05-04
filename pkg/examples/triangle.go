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
	}
	vgo.Init(setup, draw)
}
