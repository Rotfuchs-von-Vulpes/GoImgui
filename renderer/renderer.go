package renderer

import (
	"GoImgui/util"
	_ "embed"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const FLOAT_SIZE = 4

//go:embed shaders/triangle.vert
var vertexShader string

//go:embed shaders/triangle.frag
var fragmentShader string

type uniforms struct {
	color int32
}

type renderer struct {
	shaderHandle uint32
	VAO          uint32
	VBO          uint32
	uniforms     uniforms
}

var r renderer

func glError(handle uint32, statusType uint32, getIV func(uint32, uint32, *int32), getInfoLog func(uint32, int32, *int32, *uint8), failureMsg string) {
	var status int32
	getIV(handle, statusType, &status)
	if status == gl.FALSE {
		var logLength int32
		getIV(handle, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength))
		getInfoLog(handle, logLength, nil, gl.Str(infoLog))
		fmt.Println(failureMsg+"\n", infoLog)
	}
}

func Init() {
	glShaderSource := func(handle uint32, source string) {
		csource, free := gl.Strs(source + "\x00")
		defer free()

		gl.ShaderSource(handle, 1, csource, nil)
	}

	r.shaderHandle = gl.CreateProgram()
	vertHandle := gl.CreateShader(gl.VERTEX_SHADER)
	fragHandle := gl.CreateShader(gl.FRAGMENT_SHADER)
	glShaderSource(vertHandle, vertexShader)
	glShaderSource(fragHandle, fragmentShader)
	gl.CompileShader(vertHandle)
	glError(vertHandle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Vertex shader error")
	gl.CompileShader(fragHandle)
	glError(vertHandle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Fragment shader error")
	gl.AttachShader(r.shaderHandle, vertHandle)
	gl.AttachShader(r.shaderHandle, fragHandle)
	gl.LinkProgram(r.shaderHandle)
	glError(r.shaderHandle, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Linking program error")
	gl.DeleteShader(vertHandle)
	gl.DeleteShader(fragHandle)

	r.uniforms.color = gl.GetUniformLocation(r.shaderHandle, util.Str("color"))

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	gl.GenVertexArrays(1, &r.VAO)
	gl.GenBuffers(1, &r.VBO)
	gl.BindVertexArray(r.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, int(FLOAT_SIZE)*len(vertices), gl.Ptr(&vertices[0]), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*FLOAT_SIZE, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

type Data interface {
}

func Render(data util.Data) {
	gl.ClearColor(data.ClearColor[0], data.ClearColor[1], data.ClearColor[2], 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(r.shaderHandle)
	gl.Uniform3f(r.uniforms.color, data.ObjectColor[0], data.ObjectColor[1], data.ObjectColor[2])
	gl.BindVertexArray(r.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func Nuke() {
	gl.UseProgram(r.shaderHandle)
	gl.DeleteVertexArrays(1, &r.VAO)
	gl.DeleteBuffers(1, &r.VBO)
	gl.DeleteProgram(r.shaderHandle)
}
