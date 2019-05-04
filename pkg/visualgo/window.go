package visualgo

import (
	"fmt"
	"runtime"

	// NOTE: github.com/go-gl/gl/v2.1/gl can be used for greater compatibility on legacy devices with older GPUs

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var loopCode func() = nil

func (w *WindowConfig) Init(setup func(), loop func()) {
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

	glfw.WindowHint(glfw.Resizable, glfw.False)

	window.MakeContextCurrent()

	// Initialize OpenGL

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println(version)

	program := gl.CreateProgram() // initializes program reference to store our shaders
	gl.LinkProgram(program)       // link program to window

	// Run setup code passed by user
	if setup != nil {
		setup()
	}

	for !window.ShouldClose() {
		// LOOP CODE HERE
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		if loop != nil {
			loop()
		}
		//loopCode()
		glfw.PollEvents()
		window.SwapBuffers()
	}
}

// func (w *WindowConfig) Loop(f func()) {
// 	loopCode = f
// 	fmt.Println("LOOP PASS THRU")
// }

// func Render(w *glfw.Window, prog uint32) {
// 	for !w.ShouldClose() {
// 		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
// 		gl.UseProgram(prog)

// 		glfw.PollEvents()
// 		w.SwapBuffers()
// 	}
// }
