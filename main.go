package main

import (
	"fmt"

	imgui "OpenGLTest/imgui/cmd"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var winTitle string = "Go-SDL2 + Go-GL"
	var winWidth, winHeight int32 = 800, 600
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	// gui := imgui.Get()
	imgui.Run()

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.SRGB_ALPHA)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, int32(winWidth), int32(winHeight))

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State == 1 && t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			case *sdl.MouseMotionEvent:
				window.SetTitle(fmt.Sprintf("MouseMotion x:%d y:%d xrel:%d yrel:%d\n", t.X, t.Y, t.XRel, t.YRel))
			}
		}
		drawgl()
		window.GLSwap()
		// gui()
	}
}

func drawgl() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(0.5, -0.5)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(-0.5, -0.5)
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex2f(0.0, 0.5)
	gl.End()
}
