package imgui

import (
	"fmt"
	"os"

	"github.com/Rotfuchs-von-Vulpes/imgui-go/v4"

	app "GoImgui/imgui/internal/input"
	"GoImgui/imgui/internal/platform"
	"GoImgui/imgui/internal/renderer"
	"GoImgui/util"
)

type Gui struct {
	context  *imgui.Context
	platform *platform.SDL
	renderer *renderer.OpenGL3
	io       *imgui.IO
	app      app.AppData
}

var g Gui

func Init() {
	g.context = imgui.CreateContext(nil)
	io := imgui.CurrentIO()

	var err error
	g.platform, err = platform.NewSDL(io, platform.SDLClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	g.renderer, err = renderer.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	app.Init(g.platform)
}

// func (s *Gui) Run() {
// 	s.context.SetCurrent()
// 	example.Run(s.platform, s.renderer)
// }

func PreRender() {
	app.PreRender(g.platform, g.renderer)
}

func Render() {
	app.Render(g.platform, g.renderer)
}

func GetData() util.Data {
	return app.GetData()
}

func ShouldStop() bool {
	return g.platform.ShouldStop()
}

func Close() {
	g.context.SetCurrent()
	g.renderer.Dispose()
	g.platform.Dispose()
	g.context.Destroy()
}
