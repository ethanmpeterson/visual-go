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

	vgo.Setup(func() {
		fmt.Println("PASSED CODE RAN")
	})
}
