package glrender

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth = 800
const windowHeight = 600

var vertexShader = `
#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 330
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
    outputColor = vec4(fragTexCoord,0.5,0);
}
` + "\x00"

type Render struct {
	Components []Component
	Camera     Camera
	Window     *glfw.Window
	program    uint32
}

func NewRender() (*Render, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cube", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		return nil, err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	render := &Render{
		Window:  window,
		program: program,
	}
	render.Camera = NewBaseCamera(render)
	return render, nil

}

func (render *Render) Close() {
	glfw.Terminate()
}

func (render *Render) Run() {
	for !render.Window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(render.program)
		//更新相机
		render.Camera.Update()
		render.Camera.Updated()
		//更新组件
		for _, component := range render.Components {
			component.Update()
			component.Updated()
			gl.BindVertexArray(component.getVao())
			gl.DrawArrays(gl.TRIANGLES, 0, component.size())
		}
		// Maintenance
		render.Window.SwapBuffers()
		glfw.PollEvents()
	}
}
