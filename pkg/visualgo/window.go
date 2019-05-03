package visualgo

import (
	"runtime"

	// OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
)

func (w *WindowConfig) CreateWindow() {
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
	//version := gl.GoStr(gl.GetString(gl.VERSION))
	//fmt.Println(version)

}
