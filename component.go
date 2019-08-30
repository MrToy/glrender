package glrender

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Component interface {
	Update()
	updated()
	getModel() mgl32.Mat4
	size() int32
	getVao() uint32
}

type BaseComponent struct {
	Component
	Vao      uint32
	Model    mgl32.Mat4
	Vertices []float32
	Render   *Render
}

func NewBaseComponent(render *Render, vertices []float32) *BaseComponent {
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(render.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(render.program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
	return &BaseComponent{
		Vao:      vao,
		Model:    mgl32.Ident4(),
		Vertices: vertices,
		Render:   render,
	}
}

func (component *BaseComponent) Update() {

}

func (component *BaseComponent) updated() {
	modelUniform := gl.GetUniformLocation(component.Render.program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &component.Model[0])
}

func (component *BaseComponent) size() int32 {
	return int32(len(component.Vertices) / 5 * 3)
}

func (component *BaseComponent) getModel() mgl32.Mat4 {
	return component.Model
}

func (component *BaseComponent) getVao() uint32 {
	return component.Vao
}
