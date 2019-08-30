package glrender

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera interface {
	Update()
	Updated()
}

type BaseCamera struct {
	Camera
	Model      mgl32.Mat4
	Projection mgl32.Mat4
	Render     *Render
}

func NewBaseCamera(render *Render) *BaseCamera {
	return &BaseCamera{
		Render:     render,
		Model:      mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0}),
		Projection: mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 10.0),
	}
}

func (camera *BaseCamera) Update() {

}

func (camera *BaseCamera) Updated() {
	cameraUniform := gl.GetUniformLocation(camera.Render.program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera.Model[0])
	projectionUniform := gl.GetUniformLocation(camera.Render.program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &camera.Projection[0])
}
