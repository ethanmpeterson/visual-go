package visualgo

import (
	"fmt"
	"runtime"

	// NOTE: github.com/go-gl/gl/v2.1/gl can be used for greater compatibility on legacy devices with older GPUs

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (w *WindowConfig) CreateWindow() *glfw.Window {
	runtime.LockOSThread()

	// Initialize GLFW Library
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(w.Width, w.Height, w.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	// Initialize OpenGL

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println(version)

	program := gl.CreateProgram() // initializes program reference to store our shaders
	gl.LinkProgram(program)       // link program to window

	return window
}
