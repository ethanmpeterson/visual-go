package visualgo

import (
	"fmt"
	"runtime"
	"strings"

	// NOTE: github.com/go-gl/gl/v2.1/gl can be used for greater compatibility on legacy devices with older GPUs
	// or 4.1-core
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

	vertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram() // initializes program reference to store our shaders
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program) // link program to window

	// Make Vertex Array Object from vertices in draw.go
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertexList), gl.Ptr(vertexList), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// Run setup code passed by user
	if setup != nil {
		setup()
	}

	for !window.ShouldClose() {
		// LOOP CODE HERE
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program)

		// Triangle Test
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertexList)/3))

		if loop != nil {
			loop()
		}

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()

	gl.CompileShader(shader)

	var compileStatus int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &compileStatus)
	if compileStatus == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
