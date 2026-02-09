package main

import (
	gui "GoImgui/imgui/cmd"
	"GoImgui/renderer"
)

func main() {
	var running bool

	gui.Init()
	renderer.Init()

	running = true
	for running {
		data := gui.GetData()
		gui.PreRender()
		renderer.Render(data)
		gui.Render()
		if gui.ShouldStop() {
			running = false
		}
	}
	gui.Close()
	renderer.Nuke()
}
