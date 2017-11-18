package lab6

import (
	"github.com/vitalibaumtrok/glut/glut-master"
	"github.com/go-gl/gl/v2.1/gl"
	"unsafe"
	"runtime"
)

const (
	lightMoving = false

	ambientLightEnabled = false
	ambientLightRotating = true

	pointLightEnabled = true
	pointLightRotating = false
)

var (
	rotateX = float32(0)
	rotateY = float32(0)
	rotateZ = float32(0)

	noMat           = (*float32)(unsafe.Pointer(&[]float32 {0.0, 0.0, 0.0, 1.0}))
	matAmbientColor = (*float32)(unsafe.Pointer(&[]float32 {0.8, 0.8, 0.2, 1.0}))
	matDiffuse      = (*float32)(unsafe.Pointer(&[]float32 {0.1, 0.5, 0.8, 1.0}))
	matSpecular     = (*float32)(unsafe.Pointer(&[]float32 {1.0, 1.0, 1.0, 1.0}))
	shininess       = float32(100.0) // noShininess => 0.0; lowShininess => 5.0; highShininess => 100.0
	//matEmission     = (*float32)(unsafe.Pointer(&[]float32 {0.3, 0.2, 0.2, 0.0}))
)

func init() {
	runtime.LockOSThread()
}

func StartLab6() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	glut.Init()
	glut.InitDisplayMode(glut.DOUBLE | glut.RGB | glut.DEPTH)

	glut.InitWindowSize(800, 800)
	glut.InitWindowPosition(100, 100)
	glut.CreateWindow("Cube")

	gl.ClearColor (0.1, 0.4, 0.8, 0.8)
	gl.Enable(gl.LIGHTING)
	gl.LightModelf(gl.LIGHT_MODEL_TWO_SIDE, gl.TRUE)
	gl.Enable(gl.DEPTH_TEST)

	glut.DisplayFunc(display)
	glut.SpecialFunc(keyInput)
	glut.MainLoop()
}

func keyInput(key, x, y int) {
	if key == glut.KEY_RIGHT {
		rotateY += 0.5
	} else if key == glut.KEY_LEFT {
		rotateY -= 0.5
	} else if key == glut.KEY_UP {
		rotateX += 0.5
	} else if key == glut.KEY_DOWN {
		rotateX -= 0.5
	} else if key == glut.KEY_F1 {
		rotateZ -= 0.5
	} else if key == glut.KEY_F2 {
		rotateZ += 0.5
	}
	glut.PostRedisplay()
}

func initAmbientLight() {
	light0Diffuse := []float32{1.0, 1.0, 1.0}
	pLight0Diffuse := unsafe.Pointer(&light0Diffuse)
	light0Direction := []float32{0.0, 0.0, 1.0, 0.0}
	pLight0Direction := unsafe.Pointer(&light0Direction)

	gl.Enable(gl.LIGHT0)

	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, (*float32)(pLight0Diffuse))
	gl.Lightfv(gl.LIGHT0, gl.POSITION, (*float32)(pLight0Direction))
}

func initPointLight() {
	var light1Diffuse = float32(1.0)
	var sp = float32(1)
	var light1Position = float32(-1.6)
	/*
	var light1Diffuse = []float32{1.0, 1.0, 1.0}
	var sp = []float32{1,1,1,1}
	var light1Position = []float32{0.0, 0.0, -1.6, 1.0}
	*/

	gl.Enable(gl.LIGHT1)
	gl.Lightfv(gl.LIGHT1, gl.DIFFUSE, &light1Diffuse)
	gl.Lightfv(gl.LIGHT1, gl.SPECULAR, &sp)
	gl.Lightfv(gl.LIGHT1, gl.POSITION, &light1Position)

	gl.Lightf(gl.LIGHT1, gl.CONSTANT_ATTENUATION, 0.0)
	gl.Lightf(gl.LIGHT1, gl.LINEAR_ATTENUATION, 0.6)
	gl.Lightf(gl.LIGHT1, gl.QUADRATIC_ATTENUATION, 0.6)
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.LoadIdentity()

	if ambientLightEnabled && !ambientLightRotating {
		initAmbientLight()
	}
	if pointLightEnabled && !pointLightRotating {
		initPointLight()
	}

	gl.Rotatef(rotateX, 1.0, 0.0, 0.0 )
	gl.Rotatef(rotateY, 0.0, 1.0, 0.0 )
	gl.Rotatef(rotateZ, 0.0, 0.0, 1.0 )

	if ambientLightEnabled && ambientLightRotating {
		initAmbientLight()
	}
	if pointLightEnabled && pointLightRotating {
		initPointLight()
	}

	if lightMoving {
		gl.LoadIdentity()
	}

	drawConstY(0.5, matDiffuse)
	drawConstY(-0.5, matDiffuse)

	drawConstZ(0.5, matDiffuse)
	drawConstZ(-0.5, matDiffuse)

	drawConstX(0.5, matDiffuse)
	drawConstX(-0.5, matDiffuse)


	gl.Disable(gl.LIGHT0)
	gl.Disable(gl.LIGHT1)

	gl.Flush()
	glut.SwapBuffers()
}

func drawConstZ(z float32, color *float32) {
	gl.Begin(gl.QUADS)
	var x, y float32

	if z > 0 {
		gl.Normal3f(0.0, 0.0, 1.0)
	} else {
		gl.Normal3f(0.0, 0.0, -1.0)
	}

	for x = -0.5; x < 0.5; x += 0.005 {
		for y = -0.5; y < 0.5; y += 0.005 {
			if z > 0 {
				gl.Normal3f(0.0, 0.0, 1.0)
			} else {
				gl.Normal3f(0.0, 0.0, -1.0)
			}
			gl.Materialfv(gl.FRONT, gl.AMBIENT, matAmbientColor)
			gl.Materialfv(gl.FRONT, gl.DIFFUSE, color)
			gl.Materialfv(gl.FRONT, gl.SPECULAR, matSpecular)
			gl.Materialf(gl.FRONT, gl.SHININESS, shininess)
			gl.Materialfv(gl.FRONT, gl.EMISSION, noMat)
			gl.Vertex3f(x, y, z)
			gl.Vertex3f(x, y + 0.005, z)
			gl.Vertex3f(x + 0.005, y + 0.005, z)
			gl.Vertex3f(x + 0.005, y, z)
		}
	}
	gl.End()
}

func drawConstX(x float32, color *float32) {
	gl.Begin(gl.QUADS)
	var z, y float32

	if x > 0 {
	gl.Normal3f(1.0, 0.0, 0.0)
	} else {
	gl.Normal3f(-1.0, 0.0, 0.0)
	}

	for z = -0.5; z < 0.5; z += 0.005 {
		for y = -0.5; y < 0.5; y += 0.005 {
			if x > 0 {
				gl.Normal3f(1.0, 0.0, 0.0)
			} else {
				gl.Normal3f(-1.0, 0.0, 0.0)
			}
			gl.Materialfv(gl.FRONT, gl.AMBIENT, matAmbientColor)
			gl.Materialfv(gl.FRONT, gl.DIFFUSE, color)
			gl.Materialfv(gl.FRONT, gl.SPECULAR, matSpecular)
			gl.Materialf(gl.FRONT, gl.SHININESS, shininess)
			gl.Materialfv(gl.FRONT, gl.EMISSION, noMat)
			gl.Vertex3f(x, y, z)
			gl.Vertex3f(x, y + 0.005, z)
			gl.Vertex3f(x, y + 0.005, z + 0.005)
			gl.Vertex3f(x, y, z + 0.005)
		}
	}
	gl.End()
}

func drawConstY(y float32, color *float32) {
	gl.Begin(gl.QUADS)
	var x, z float32

	if y > 0 {
		gl.Normal3f(0.0, 1.0, 0.0)
	} else {
		gl.Normal3f(0.0, -1.0, 0.0)
	}

	for x = -0.5; x < 0.5; x += 0.005 {
		for z = -0.5; z < 0.5; z += 0.005 {
			if y > 0 {
				gl.Normal3f(0.0, 1.0, 0.0)
			} else {
				gl.Normal3f(0.0, -1.0, 0.0)
			}
			gl.Materialfv(gl.FRONT, gl.AMBIENT, matAmbientColor)
			gl.Materialfv(gl.FRONT, gl.DIFFUSE, color)
			gl.Materialfv(gl.FRONT, gl.SPECULAR, matSpecular)
			gl.Materialf(gl.FRONT, gl.SHININESS, shininess)
			gl.Materialfv(gl.FRONT, gl.EMISSION, noMat)
			gl.Vertex3f(x, y, z)
			gl.Vertex3f(x, y , z + 0.005)
			gl.Vertex3f(x + 0.005, y, z + 0.005)
			gl.Vertex3f(x + 0.005, y, z)
		}
	}
	gl.End()
}
