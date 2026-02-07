package main

import (
	imgui "OpenGLImguiGoTest/imgui/cmd"
)

func main() {
	var running bool

	gui := imgui.GetImgui()
	gui.Init()

	running = true
	for running {
		gui.Render()
		if gui.ShouldStop() {
			running = false
		}
	}
	gui.Close()
}
