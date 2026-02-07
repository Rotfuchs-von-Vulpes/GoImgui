package imgui

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v4"

	"OpenGLImguiGoTest/imgui/internal/example"
	"OpenGLImguiGoTest/imgui/internal/platform"
	"OpenGLImguiGoTest/imgui/internal/renderer"
)

type Gui struct {
	context  *imgui.Context
	platform *platform.SDL
	renderer *renderer.OpenGL3
	io       *imgui.IO
	app      example.AppData
}

func GetImgui() Gui {
	return Gui{}
}

func (s *Gui) Init() {
	s.context = imgui.CreateContext(nil)
	io := imgui.CurrentIO()

	var err error
	s.platform, err = platform.NewSDL(io, platform.SDLClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	s.renderer, err = renderer.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	s.app = example.GetApp()
	s.app.Init(s.platform)
}

// func (s *Gui) Run() {
// 	s.context.SetCurrent()
// 	example.Run(s.platform, s.renderer)
// }

func (s *Gui) Render() {
	s.context.SetCurrent()
	s.app.Render(s.platform, s.renderer)
}

func (s *Gui) ShouldStop() bool {
	return s.platform.ShouldStop()
}

func (s *Gui) Close() {
	s.context.SetCurrent()
	s.renderer.Dispose()
	s.platform.Dispose()
	s.context.Destroy()
}
