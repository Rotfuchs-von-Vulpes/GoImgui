package input

import (
	"fmt"
	"time"

	"GoImgui/imgui/internal/demo"
	"GoImgui/util"

	"github.com/inkyblackness/imgui-go/v4"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type clipboard struct {
	platform Platform
}

func (board clipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board clipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

const (
	millisPerSecond = 1000
	sleepDuration   = time.Millisecond * 25
)

type AppData struct {
	showDemoWindow    bool
	showGoDemoWindow  bool
	clearColor        [3]float32
	objectColor       [3]float32
	f                 float32
	counter           int
	showAnotherWindow bool
}

var app AppData

func Init(p Platform) {
	imgui.CurrentIO().SetClipboard(clipboard{platform: p})
	app.showDemoWindow = false
	app.showGoDemoWindow = false
	app.clearColor = [3]float32{0.0, 0.0, 0.0}
	app.objectColor = [3]float32{1.0, 0.5, 0.2}
	app.f = float32(0)
	app.counter = 0
	app.showAnotherWindow = false
}

func PreRender(p Platform, r Renderer) {
	p.ProcessEvents()

	// Signal start of a new frame
	p.NewFrame()
	imgui.NewFrame()

	// 1. Show a simple window.
	// Tip: if we don't call imgui.Begin()/imgui.End() the widgets automatically appears in a window called "Debug".
	{
		imgui.Text("ภาษาไทย测试조선말")                         // To display these, you'll need to register a compatible font
		imgui.Text("Hello, world!")                        // Display some text
		imgui.SliderFloat("float", &app.f, 0.0, 1.0)       // Edit 1 float using a slider from 0.0f to 1.0f
		imgui.ColorEdit3("clear color", &app.clearColor)   // Edit 3 floats representing a color
		imgui.ColorEdit3("object color", &app.objectColor) // Edit 3 floats representing a color

		imgui.Checkbox("Demo Window", &app.showDemoWindow) // Edit bools storing our window open/close state
		imgui.Checkbox("Go Demo Window", &app.showGoDemoWindow)
		imgui.Checkbox("Another Window", &app.showAnotherWindow)

		if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
			app.counter++
		}
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("counter = %d", app.counter))

		imgui.Text(fmt.Sprintf("Application average %.3f ms/frame (%.1f FPS)",
			millisPerSecond/imgui.CurrentIO().Framerate(), imgui.CurrentIO().Framerate()))
	}

	// 2. Show another simple window. In most cases you will use an explicit Begin/End pair to name your windows.
	if app.showAnotherWindow {
		// Pass a pointer to our bool variable (the window will have a closing button that will clear the bool when clicked)
		imgui.BeginV("Another window", &app.showAnotherWindow, 0)
		imgui.Text("Hello from another window!")
		if imgui.Button("Close Me") {
			app.showAnotherWindow = false
		}
		imgui.End()
	}

	// 3. Show the ImGui demo window. Most of the sample code is in imgui.ShowDemoWindow().
	// Read its code to learn more about Dear ImGui!
	if app.showDemoWindow {
		// Normally user code doesn't need/want to call this because positions are saved in .ini file anyway.
		// Here we just want to make the demo initial state a bit more friendly!
		const demoX = 650
		const demoY = 20
		imgui.SetNextWindowPosV(imgui.Vec2{X: demoX, Y: demoY}, imgui.ConditionFirstUseEver, imgui.Vec2{})

		imgui.ShowDemoWindow(&app.showDemoWindow)
	}
	if app.showGoDemoWindow {
		demo.Show(&app.showGoDemoWindow)
	}

	// Rendering
	imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done in Render function in this package.
}

func Render(p Platform, r Renderer) {
	r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
	p.PostRender()
}

func GetData() util.Data {
	return util.GetData(app.clearColor, app.objectColor)
}
