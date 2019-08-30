package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mrtoy/glrender"
)

func init() {
	runtime.LockOSThread()
}

type Cube struct {
	*glrender.BaseComponent
	previousTime float64
}

func NewCube(render *glrender.Render) *Cube {
	cube := &Cube{
		BaseComponent: glrender.NewBaseComponent(render, cubeVertices),
		previousTime:  glfw.GetTime(),
	}
	cube.Model = mgl32.Translate3D(0, 0, 0)
	return cube
}

func (cube *Cube) Update() {
	if cube.Render.Window.GetKey(glfw.KeyD) == glfw.Press {
		cube.Model = mgl32.Translate3D(0.1, 0, 0).Mul4(cube.Model)
	}
	if cube.Render.Window.GetKey(glfw.KeyA) == glfw.Press {
		cube.Model = mgl32.Translate3D(-0.1, 0, 0).Mul4(cube.Model)
	}
	if cube.Render.Window.GetKey(glfw.KeyW) == glfw.Press {
		cube.Model = mgl32.Translate3D(0, 0, -0.1).Mul4(cube.Model)
	}
	if cube.Render.Window.GetKey(glfw.KeyS) == glfw.Press {
		cube.Model = mgl32.Translate3D(0, 0, 0.1).Mul4(cube.Model)
	}
	time := glfw.GetTime()
	elapsed := time - cube.previousTime
	cube.previousTime = time
	cube.Model = cube.Model.Mul4(mgl32.HomogRotate3D(float32(elapsed), mgl32.Vec3{0, 1, 0}))
}

type Cube2 struct {
	*glrender.BaseComponent
	previousTime float64
}

func NewCube2(render *glrender.Render) *Cube2 {
	cube := &Cube2{
		BaseComponent: glrender.NewBaseComponent(render, cubeVertices),
		previousTime:  glfw.GetTime(),
	}
	cube.Model = mgl32.Translate3D(2, 2, 0)
	return cube
}

func (cube *Cube2) Update() {
	time := glfw.GetTime()
	elapsed := time - cube.previousTime
	cube.previousTime = time
	cube.Model = cube.Model.Mul4(mgl32.HomogRotate3D(float32(-elapsed), mgl32.Vec3{0, 1, 0}))
}

type Camera struct {
	*glrender.BaseCamera
}

const windowWidth = 800
const windowHeight = 600

func NewCamera(render *glrender.Render) *Camera {
	camera := &Camera{
		BaseCamera: glrender.NewBaseCamera(render),
	}
	camera.Model = mgl32.LookAtV(mgl32.Vec3{10, 10, 10}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	camera.Projection = mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 100.0)
	return camera
}

func (camera *Camera) onScroll(w *glfw.Window, xoff float64, yoff float64) {
	camera.Model = mgl32.Translate3D(0, 0, float32(yoff/10)).Mul4(camera.Model)
}

func (camera *Camera) Update() {
	if camera.Render.Window.GetKey(glfw.KeyRight) == glfw.Press {
		camera.Model = mgl32.Translate3D(0.1, 0, 0).Mul4(camera.Model)
	}
	if camera.Render.Window.GetKey(glfw.KeyLeft) == glfw.Press {
		camera.Model = mgl32.Translate3D(-0.1, 0, 0).Mul4(camera.Model)
	}
	if camera.Render.Window.GetKey(glfw.KeyUp) == glfw.Press {
		camera.Model = mgl32.Translate3D(0, -0.1, 0).Mul4(camera.Model)
	}
	if camera.Render.Window.GetKey(glfw.KeyDown) == glfw.Press {
		camera.Model = mgl32.Translate3D(0, 0.1, 0).Mul4(camera.Model)
	}
}

func main() {
	render, err := glrender.NewRender()
	if err != nil {
		panic(err)
	}
	defer render.Close()
	cemara := NewCamera(render)
	render.Camera = cemara
	render.Components = []glrender.Component{
		NewCube(render),
		NewCube2(render),
	}
	render.Window.SetScrollCallback(func(w *glfw.Window, xoff float64, yoff float64) {
		cemara.onScroll(w, xoff, yoff)
	})
	render.Run()
}

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}
