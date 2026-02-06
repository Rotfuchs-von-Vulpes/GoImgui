package imgui

import (
	"fmt"
	"os"

	"github.com/inkyblackness/imgui-go/v4"

	"OpenGLTest/imgui/internal/example"
	"OpenGLTest/imgui/internal/platform"
	"OpenGLTest/imgui/internal/renderer"
)

func Get() func() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platform.NewSDL(io, platform.SDLClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderer.NewOpenGL2(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	return example.GetRenderFunc(platform, renderer)
}

func Run() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platform.NewSDL(io, platform.SDLClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderer.NewOpenGL2(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	example.Run(platform, renderer)
}
